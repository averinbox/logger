package main

import "github.com/averinbox/logger"

func main() {
	// Enable dev mode
	log := logger.New(true)

	log.Info("Info message ...")
	log.Warn("Warn message ...")
	log.Debug("Debug message ...")
	log.Error("Error message ...")
}
