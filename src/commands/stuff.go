package commands

import (
	"crypto/rand"
	"math/big"
	"regexp"
	"roadstatebot/src/bot"
)

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
			"http://risovach.ru/upload/2020/07/mem/hitriy-getsbi_244363555_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/v-mire-grustit-odin-kotik_244363635_orig_.jpg",
			"http://risovach.ru/upload/2020/07/mem/velikiy-getsbi-bokal-za-teh_244363697_orig_.jpg",
		}

		return &bot.Message{PhotoURL: getRandValueInArr(imgs)}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^(300|триста)$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		imgs := []string{
			"https://pbs.twimg.com/media/EcGH3k-XQAABbRQ?format=jpg",
			"https://www.meme-arsenal.com/memes/0e0fb895ce30bc291f39a928d5d9ff6c.jpg",
			"https://www.meme-arsenal.com/memes/f14353ecdbc9ada0d08cb4c0f98f6083.jpg",
			"https://www.meme-arsenal.com/memes/8de3bb6fabaa7b86146bde0397f618f7.jpg",
			"https://www.meme-arsenal.com/memes/f446a4b4c17a220a128de1cd7153c03b.jpg",
		}

		return &bot.Message{PhotoURL: getRandValueInArr(imgs)}
	})

	ibot.OnRegexp(regexp.MustCompile(`(?i)^нет$`), func(*bot.User, *bot.Chat, *bot.Message) *bot.Message {
		imgs := []string{
			"https://www.meme-arsenal.com/memes/8d030eb87804939de063b97cb0ec5439.jpg",
			"https://www.meme-arsenal.com/memes/dffd4b48d57fd0b2b95aa3aa240cbea8.jpg",
			"https://www.meme-arsenal.com/memes/2f3066597c524b2aa6f6a8e36ebd0b7b.jpg",
		}

		return &bot.Message{PhotoURL: getRandValueInArr(imgs)}
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