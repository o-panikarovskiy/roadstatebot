package commands

import (
	"crypto/rand"
	"math/big"
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
	ibot.OnAnswer("/highways", HighwayListAnswer(rep))

	ibot.OnRegexp(regexp.MustCompile(`(?i)^/(ua|ru|by|uа)([MPAHМРАН])[0-9]+$`), FeedbacksList(rep))
	ibot.OnAnswer("/feedback", FeedbackAnswer(rep))

	initStuff(ibot)
}

func initStuff(ibot bot.IBot) {
	ibot.OnRegexp(regexp.MustCompile(`(?i)^да$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		text := "пизда"

		b, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err == nil && b.Int64()%2 == 0 {
			text = "хуй на"
		}

		return &bot.Message{Text: text}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^нет$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		text := "пидара ответ"

		b, err := rand.Int(rand.Reader, big.NewInt(1000))
		if err == nil && b.Int64()%2 == 0 {
			text = "шлюхи аргумент"
		}

		return &bot.Message{Text: text}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^(макс|максим|максимка|антон|тоха|виталий|веталь|витаха|олег|олежа)$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		return &bot.Message{Text: "ты красавчик"}
	})
}
