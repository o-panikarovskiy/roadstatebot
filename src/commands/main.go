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
		imgs := []string{
			"https://pbs.twimg.com/media/EGd7J1bXUAAz_-E.jpg",
			"http://risovach.ru/upload/2020/07/mem/ilon-mask_244328397_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/otchayannyy-agutin_244328442_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/diana-shurygina_244328487_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/ne-delay-ne-budet_244328550_orig_.jpg",
			"http://risovach.ru/upload/2016/12/mem/penapple_131242655_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/harley-quinn_244330188_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/boloto_244330206_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/bezlimiticshe_244330235_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/toni-stark_244330259_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/zheleznyy-chelovek_244330279_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/chuvak-eto-repchik_244330300_orig_.jpg",
			"https://pbs.twimg.com/media/EbefYrsWsAAIeFL?format=jpg&name=medium",
			"https://pbs.twimg.com/media/Ebef4oSXYAAtUWH?format=jpg&name=medium",
		}

		return &bot.Message{PhotoURL: getRandValueInArr(imgs)}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^(300|триста)$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		imgs := []string{
			"https://pbs.twimg.com/media/EcGH3k-XQAABbRQ?format=jpg",
		}

		return &bot.Message{PhotoURL: getRandValueInArr(imgs)}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^нет$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		return &bot.Message{Text: "пидара ответ"}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^(макс|максим|максимка)$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		return &bot.Message{Text: "ты плоечку купил?"}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^(антон|тоха|виталий|веталь|витаха|олег|олежа)$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		return &bot.Message{Text: "ты красавчик"}
	})
}

func getRandValueInArr(arr []string) string {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(arr))))

	var idx int64
	if err == nil {
		idx = index.Int64()
	} else {
		idx = 0
	}

	return arr[idx]
}
