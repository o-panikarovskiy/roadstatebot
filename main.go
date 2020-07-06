package main

import (
	"log"
	"os"
	"os/signal"

	"roadstatebot/src/server"
)

func main() {
	botAPIKey := os.Getenv("TG_BOT_KEY")
	if botAPIKey == "" {
		log.Println("Please, specify TG_BOT_KEY")
		return
	}

	instance := server.NewInstance(botAPIKey)
	instance.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	instance.Stop()
	os.Exit(0)
}
