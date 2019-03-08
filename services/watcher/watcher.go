package watcher

import "watchmen/services"

type Options struct {
	Company string `json:"company"`
}

type Option func(option *Options)

func WithCompany(company string) Option {
	return func(options *Options) {
		options.Company = company
	}
}

func CheckIn(user string, password string, watcherOptions ...Option) Options {
	options := Options{Company: "42"}
	for _, opt := range watcherOptions {
		opt(&options)
	}
	services.Logger.Info("OPTIONS", options)
	return options
}
