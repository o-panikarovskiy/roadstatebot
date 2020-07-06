package commands

import (
	"regexp"
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
)

// Init command
func Init(ibot bot.IBot, rep repository.IRepository) {
	botName := ibot.BotName()

	ibot.On("/start", Start())
	ibot.On("/restart", Start())
	ibot.On("/help", Help())

	ibot.On("/list", CountriesList(rep))
	ibot.On("/list@"+botName, CountriesList(rep))
	ibot.OnAnswer("/highways", HighwayListAnswer(rep))

	ibot.OnRegexp(regexp.MustCompile(`(?i)^/(ua|ru|by|uа)([MPAHМРАН])[0-9]+$`), FeedbacksList(rep))
	ibot.OnAnswer("/feedback", FeedbackAnswer(rep))

	ibot.OnRegexp(regexp.MustCompile("^!"), AnfisaChat)
}
