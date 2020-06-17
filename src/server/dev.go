package server

import (
	"log"
	"regexp"
	"roadstatebot/src/bot"
	"roadstatebot/src/commands"
	"roadstatebot/src/config"
	"roadstatebot/src/repository/autostrada"
)

func createDevInstase(cfg *config.AppConfig) *Instance {
	bot, err := bot.New(cfg)

	if err != nil {
		log.Panic(err)
	}

	botName := bot.BotName()
	log.Printf("Authorized on account %s", bot.BotName())

	repository := autostrada.New()

	bot.On("/start", commands.Start(cfg))
	bot.On("/restart", commands.Start(cfg))
	bot.On("/help", commands.Help(cfg))

	bot.On("/list", commands.Counties(repository))
	bot.On("/list@"+botName, commands.Counties(repository))
	bot.OnAnswer("/list", commands.CountiesAnswer(repository))
	bot.OnAnswer("/highway", commands.HighwayTypeAnswer(repository))

	regHighway, _ := regexp.Compile(`(?i)^/(ua|ru|by|uа)([MPAHМРАН])[0-9]+$`)
	bot.OnRegexp(regHighway, commands.Feedback(repository))
	bot.OnAnswer("/feedback", commands.FeedbackAnswer(repository))

	return &Instance{
		bot: bot,
		cfg: cfg,
	}
}
