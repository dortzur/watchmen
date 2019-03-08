package watcher

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty"
	"math/rand"
	"strconv"
	"time"
)

type Options struct {
	Company string `json:"company"`
}

var defaultOptions = Options{Company: "4266"}

type Option func(option *Options)

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

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
	client.SetRedirectPolicy(resty.FlexibleRedirectPolicy(15))
	response, err := client.R().SetFormData(map[string]string{
		"comp": options.Company,
		"name": user,
		"pw":   password,
		"B1.x": strconv.Itoa(random.Intn(30-1) + 1),
		"B1.y": strconv.Itoa(random.Intn(30-1) + 1),
	}).Post("http://checkin.timewatch.co.il/punch/punch2.php")

	headers := response.Header()
	if err != nil {
		return nil, err
	}
	if headers.Get("Set-Cookie") == "" {
		return nil, errors.New("authentication failed. Please check credentials")
	}

	return gin.H{"user": user, "password": password, "options": options}, nil
}

func CheckOut(user string, password string, watcherOptions ...Option) Options {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
	return options
}
