package main

import (
	"io"
	"log"
)

var (
	// Info wites logs at the "info" level
	Info *log.Logger
	// Debug writes logs at the "debug" level
	Debug *log.Logger
	// Error writes logs at the "error level"
	Error *log.Logger
)

// InitLoggers initialises the logging system
func InitLoggers(debugLogger io.Writer, infoLogger io.Writer, errorLogger io.Writer) {
	Debug = log.New(debugLogger, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoLogger, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorLogger, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}
