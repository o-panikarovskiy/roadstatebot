package server

import (
	"log"
	"roadstatebot/src/bot"
	"roadstatebot/src/commands"
	"roadstatebot/src/repository/autostrada"
)

func createDevInstase(apiKey string) *Instance {
	ibot, err := bot.New(apiKey)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", ibot.BotName())
	commands.Init(ibot, autostrada.New())

	return &Instance{bot: ibot}
}
