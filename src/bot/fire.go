package bot

import (
	"fmt"
	"net/url"

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

	if message.File != nil {
		tgFileMsg := tgbotapi.NewPhotoUpload(update.Message.Chat.ID, inst.toTgFile(message.File))
		if !update.Message.Chat.IsPrivate() {
			tgFileMsg.ReplyToMessageID = update.Message.MessageID
		}
		inst.api.Send(tgFileMsg)
		return
	}

	if message.PhotoURL != "" {
		params := url.Values{}
		params.Add("chat_id", fmt.Sprint(update.Message.Chat.ID))
		params.Add("photo", message.PhotoURL)
		params.Add("caption", message.Text)
		if !update.Message.Chat.IsPrivate() {
			params.Add("reply_to_message_id", fmt.Sprint(update.Message.MessageID))
		}

		inst.api.MakeRequest("sendPhoto", params)
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
