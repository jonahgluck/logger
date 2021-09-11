package main

import (
	// log "github.com/jonahgluck/logger"
	log "logger/log"
)

var settings = []string{
	// "bold",
	"none",
}

func main() {
    log.Log("Warning:", "magenta", settings)
}
