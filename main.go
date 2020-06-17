package main

// import (
// 	"log"

// 	"github.com/go-telegram-bot-api/telegram-bot-api"
// )

// func main() {
// 	bot, err := tgbotapi.NewBotAPI("MyAwesomeBotToken")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	bot.Debug = true

// 	log.Printf("Authorized on account %s", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates, err := bot.GetUpdatesChan(u)

// 	for update := range updates {
// 		if update.Message == nil { // ignore any non-Message Updates
// 			continue
// 		}

// 		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 		msg.ReplyToMessageID = update.Message.MessageID

// 		bot.Send(msg)
// 	}
// }

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"

	"roadstatebot/src/config"
	"roadstatebot/src/server"
)

func main() {
	instance := server.NewInstance(setup())

	instance.Run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT (Ctrl+C)
	<-c                            // Block until we receive our signal.

	instance.Stop()
	os.Exit(0)
}

func setup() *config.AppConfig {
	if len(os.Args) < 2 {
		log.Panicln("Please, specify the config file")
	}

	path, err := filepath.Abs(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}

	return config.NewDefaultConfig(path)
}
