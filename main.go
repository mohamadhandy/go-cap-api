package main

import (
	"capi/app"
	"capi/logger"
	"log"
)

func main() {
	log.Println("Starting Application")
	logger.Info("Starting application")
	app.Start()
}
