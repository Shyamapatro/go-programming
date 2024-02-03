package main

import (
    "log"
    "os"
	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {
     // Create a new logger
    logger := log.New(os.Stdout, "", 0)

    // Create colorized loggers
    infoLogger := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
    warnLogger := color.New(color.FgYellow).Add(color.Bold).SprintFunc()
    errorLogger := color.New(color.FgRed).Add(color.Bold).SprintFunc()

    // Log messages with colors
    logger.Println("This is a", infoLogger("info"), "message.")
    logger.Println("This is a", warnLogger("warning"), "message.")
    logger.Println("This is an", errorLogger("error"), "message.")
	println(uuid.New().String())
}
