package server

import (
	"watchmen/services"
)

func Init() {
	//conf := config.GetConfig()
	r := NewRouter()
	err := r.Run()
	if err != nil {
		services.Logger.Fatal(err)
	}
}
