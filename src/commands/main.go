package commands

import (
	"regexp"
	"roadstatebot/src/bot"
	"roadstatebot/src/config"
	"roadstatebot/src/repository"
)

// Init command
func Init(ibot bot.IBot, cfg *config.AppConfig, rep repository.IRepository) {
	botName := ibot.BotName()

	ibot.On("/start", Start(cfg))
	ibot.On("/restart", Start(cfg))
	ibot.On("/help", Help(cfg))

	ibot.On("/list", CountriesList(rep))
	ibot.On("/list@"+botName, CountriesList(rep))
	ibot.OnAnswer("/list", CountiesListAnswer(rep))
	ibot.OnAnswer("/highways", HighwayListAnswer(rep))

	regHighway, _ := regexp.Compile(`(?i)^/(ua|ru|by|uа)([MPAHМРАН])[0-9]+$`)
	ibot.OnRegexp(regHighway, FeedbacksList(rep))
	ibot.OnAnswer("/feedback", FeedbackAnswer(rep))
}
