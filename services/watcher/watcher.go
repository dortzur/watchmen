package watcher

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

type Options struct {
	Company string `json:"company"`
}

var defaultOptions = Options{Company: "4266"}

type Option func(option *Options)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))
var ixeeRegex, _ = regexp.Compile("ixee: (.*)}")

func WithCompany(company string) Option {
	return func(options *Options) {
		if company != "" {
			options.Company = company
		}
	}
}

func CheckIn(user string, password string, watcherOptions ...Option) (map[string]interface{}, error) {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
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
	response, err := client.R().SetFormData(map[string]string{
		"comp": options.Company,
		"name": user,
		"pw":   password,
		"B1.x": strconv.Itoa(random.Intn(30-1) + 1),
		"B1.y": strconv.Itoa(random.Intn(30-1) + 1),
	}).Post("http://checkin.timewatch.co.il/punch/punch2.php")
	if err != nil {
		return nil, err
	}
	cookies := response.Cookies()

	if len(cookies) == 0 {
		return nil, errors.New("authentication failed. Please check credentials")
	}
	ixeeResult := ixeeRegex.FindStringSubmatch(response.String())
	if len(ixeeResult) < 2 {
		return nil, errors.New("couldn't parse timewatch response")
	}
	ixee := ixeeResult[1]
	client.Cookies = cookies
	client.SetDebug(true)

	layout := "2006-01-02 15:04:05"
	now := time.Now().Format(layout)
	response, err = client.R().SetFormData(map[string]string{
		"comp":          options.Company,
		"name":          user,
		"ts":            now,
		"ix":            ixee,
		"B1":            "כניסה",
		"allowremarks":  "1",
		"msgfound":      "0",
		"thetask":       "0",
		"teamleader":    "0",
		"speccomp":      "",
		"remark":        "",
		"tasks":         "",
		"taskdescr":     "",
		"prevtask":      "0",
		"prevtaskdescr": "",
		"withtasks":     "0",
		"tflag":         "",
	}).Post("http://checkin.timewatch.co.il/punch/punch3.php")

	//checkout params
	//B1: יציאה
	//tflag: 1

	//services.Logger.Info(response.String())

	return gin.H{"user": user, "password": password, "options": options}, nil
}

func CheckOut(user string, password string, watcherOptions ...Option) Options {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
	return options
}
