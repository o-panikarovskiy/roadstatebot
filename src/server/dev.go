package server

import (
	"log"
	"roadstatebot/src/bot"
	"roadstatebot/src/commands"
	"roadstatebot/src/config"
	"roadstatebot/src/repository/autostrada"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	ibot, err := bot.New(cfg)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", ibot.BotName())
	commands.Init(ibot, cfg, autostrada.New())

	return &Instance{
		bot: ibot,
		cfg: cfg,
	}
}
