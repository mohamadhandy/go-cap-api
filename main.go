package main

import (
	"capi/app"
	"capi/logger"
)

func main() {
	logger.Info("Starting application")
	app.Start()
}
