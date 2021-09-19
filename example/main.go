package main

import (
	// logger "github.com/jonahgluck/logger"
	logger "logger/log"
)

var settings = []string{
	"none",
}

func main() {
    logger.Log("Warning:", "magenta", settings)
    logger.Log("Warning:", "", settings)
    logger.Log("Warning:", "blue", settings)
}
