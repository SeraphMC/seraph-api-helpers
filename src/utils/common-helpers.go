package utils

import "os"

func IsDevelopmentMode() bool {
	devMode := os.Getenv("APP_ENV") == "DEV"
	return devMode
}
