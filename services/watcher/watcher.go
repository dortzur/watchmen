package watcher

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"watchmen/model"
	"watchmen/services/watcher/watcherOperation"
	"watchmen/services/watcher/watcherParams"
)

type Options struct {
	Company string `json:"company"`
}

var defaultOptions = Options{Company: "4266"}

type Option func(option *Options)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var ixeeRegex, _ = regexp.Compile("ixee: (.*)}")

const loginUrl = "https://checkin.timewatch.co.il/punch/punch2.php"
const checkinCheckoutUrl = "https://checkin.timewatch.co.il/punch/punch3.php"

func WithCompany(company string) Option {
	return func(options *Options) {
		if company != "" {
			options.Company = company
		}
	}
}

func getClient() *resty.Client {
	client := resty.New()
	client.SetHeaders(
		map[string]string{
			"Host":             "checkin.timewatch.co.il",
			"Accept":           "application/json, text/javascript, */*; q=0.01",
			"X-Requested-With": "XMLHttpRequest",
			"User-Agent":       "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
			"Referer":          "http://checkin.timewatch.co.il/punch/punch2.php",
			"Accept-Language":  "en-US,en;q=0.9,he;q=0.8,la;q=0.7",
		})

	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	return client
}

func doLogin(user string, password string, company string) (*resty.Client, string, error) {
	client := getClient()

	response, err := client.R().SetFormData(map[string]string{
		"comp": company,
		"name": user,
		"pw":   password,
		"B1.x": strconv.Itoa(random.Intn(30-1) + 1),
		"B1.y": strconv.Itoa(random.Intn(30-1) + 1),
	}).Post(loginUrl)
	if err != nil {
		return nil, "", err
	}
	cookies := response.Cookies()

	if len(cookies) == 0 {
		return nil, "", errors.New("authentication failed. Please check credentials")
	}
	ixeeResult := ixeeRegex.FindStringSubmatch(response.String())
	if len(ixeeResult) < 2 {
		return nil, "", errors.New("couldn't parse timewatch response")
	}

	ixee := ixeeResult[1]
	client.Cookies = cookies
	return client, ixee, nil
}

func doOperation(userData model.UserData, operation watcherOperation.Operation) (map[string]interface{}, error) {
	client, ixee, err := doLogin(userData.User, userData.Password, userData.Company)
	if err != nil {
		return nil, err
	}
	var params map[string]string = nil
	if operation == watcherOperation.CheckIn {
		params = watcherParams.GetCheckinParams(userData.User, userData.Company, ixee)
	} else {
		params = watcherParams.GetCheckoutParams(userData.User, userData.Company, ixee)
	}
	_, err = client.R().SetFormData(params).Post(checkinCheckoutUrl)
	return gin.H{"userData": userData, "operation": operation}, nil
}

func CheckIn(user string, password string, watcherOptions ...Option) (map[string]interface{}, error) {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
	userData := model.UserData{User: user, Password: password, Company: options.Company}
	return doOperation(userData, watcherOperation.CheckIn)
}

func CheckOut(user string, password string, watcherOptions ...Option) (map[string]interface{}, error) {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}

	userData := model.UserData{User: user, Password: password, Company: options.Company}
	return doOperation(userData, watcherOperation.CheckOut)
}
