package bot

import (
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (inst *botStruct) fire(update tgbotapi.Update) {
	handler := inst.findHandler(update.Message.Text)
	if handler == nil {
		return
	}

	msg := inst.toMessage(update.Message)
	chat := inst.toChat(update.Message.Chat)
	user := inst.toUser(update.Message.From)

	message := handler(user, chat, msg)
	if message == nil {
		return
	}

	tgMessage := tgbotapi.NewMessage(update.Message.Chat.ID, message.Text)
	tgMessage.ParseMode = "markdown"

	if !update.Message.Chat.IsPrivate() {
		tgMessage.ReplyToMessageID = update.Message.MessageID
	}

	if message.ReplyMarkup != nil {
		tgMessage.ReplyMarkup = inst.toTgInlineKeyboardMarkup(message.ReplyMarkup)
	}

	inst.api.Send(tgMessage)
}

func (inst *botStruct) fireAnswer(update tgbotapi.Update) {
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
