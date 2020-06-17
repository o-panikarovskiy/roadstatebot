package bot

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Run start listen messages
func (inst *botStruct) Run() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := inst.api.GetUpdatesChan(u)
	if err != nil {
		log.Panicln(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			go inst.fireAnswer(update)
		}

		if update.Message != nil {
			logMessage(update.Message)
			go inst.fire(update)
		}
	}
}

func logMessage(message *tgbotapi.Message) {
	var username string

	if message.From.UserName != "" {
		username = message.From.UserName
	} else {
		username = strings.TrimSpace(message.From.FirstName + " " + message.From.LastName)
	}

	log.Printf("[%s] %s", username, message.Text)
}
