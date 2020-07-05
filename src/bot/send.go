package bot

import (
	"fmt"
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (inst *botStruct) sendText(chatID int64, m *Message, replyID int) {
	tgMessage := tgbotapi.NewMessage(chatID, m.Text)
	tgMessage.ParseMode = "markdown"
	tgMessage.ReplyToMessageID = replyID

	if m.ReplyMarkup != nil {
		tgMessage.ReplyMarkup = inst.toTgInlineKeyboardMarkup(m.ReplyMarkup)
	}

	inst.api.Send(tgMessage)
}

func (inst *botStruct) sendFile(chatID int64, m *Message, replyID int) {
	tgFileMsg := tgbotapi.NewPhotoUpload(chatID, inst.toTgFile(m.File))
	tgFileMsg.ReplyToMessageID = replyID
	inst.api.Send(tgFileMsg)
}

func (inst *botStruct) sendPhotoURL(chatID int64, m *Message, replyID int) {
	params := url.Values{}
	params.Add("chat_id", fmt.Sprint(chatID))
	params.Add("photo", m.PhotoURL)
	params.Add("caption", m.Text)
	params.Add("reply_to_message_id", fmt.Sprint(replyID))

	inst.api.MakeRequest("sendPhoto", params)
}
