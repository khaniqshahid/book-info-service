package main

import (
	"github.com/khaniqshahid/book-details-service/app"
	"github.com/khaniqshahid/book-details-service/logger"
)

// go mod init github.com/ashishjuyal/BANKING

func main() {

	// Start Routes
	logger.Logger.Info("Starting the Application")
	app.Start()
}
