package services

import "os"

func GetEnv() string {
	env := os.Getenv("GO_ENV")
	if os.Getenv("GO_ENV") == "" {
		env = "development"
	}
	return env
}
