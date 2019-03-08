package watcher

import "github.com/gin-gonic/gin"

type Options struct {
	Company string `json:"company"`
}

var defaultOptions = Options{Company: "42"}

type Option func(option *Options)

func WithCompany(company string) Option {
	return func(options *Options) {
		if company != "" {
			options.Company = company
		}
	}
}

func CheckIn(user string, password string, watcherOptions ...Option) map[string]interface{} {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
	return gin.H{"user": user, "password": password, "options": options}
}

func CheckOut(user string, password string, watcherOptions ...Option) Options {
	options := defaultOptions
	for _, opt := range watcherOptions {
		opt(&options)
	}
	return options
}
