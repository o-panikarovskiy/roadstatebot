package main

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"roadstatebot/src/config"
	"roadstatebot/src/server"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please, specify the config file")
		return
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Fatal(err)
		return
	}

	instance := server.NewInstance(config.NewDefaultConfig(path))

	instance.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	instance.Stop()
	os.Exit(0)
}
