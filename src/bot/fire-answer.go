package bot

import (
	"fmt"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (inst *botStruct) fireAnswer(update tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	answerData := strings.Split(update.CallbackQuery.Data, ":")
	handler := inst.findAnswerHandler(answerData[0])
	if handler == nil {
		return
	}

	message := handler(answerData[1])
	if message == nil {
		return
	}

	tgMessage := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, message.Text)
	tgMessage.ParseMode = "markdown"

	if message.ReplyMarkup != nil {
		tgMessage.ReplyMarkup = inst.toTgInlineKeyboardMarkup(message.ReplyMarkup)
	}

	inst.api.Send(tgMessage)
}
