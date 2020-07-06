package bot

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (inst *botStruct) fire(update tgbotapi.Update) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, r)
		}
	}()

	handler := inst.findHandler(update.Message.Text)
	if handler == nil {
		return
	}

	msg := inst.toMessage(update.Message)
	chat := inst.toChat(update.Message.Chat)
	user := inst.toUser(update.Message.From)

	log.Printf("[%s %s]:%s", user.FirstName, user.LastName, msg.Text)

	message := handler(user, chat, msg)
	if message == nil {
		return
	}

	chatID := update.Message.Chat.ID
	replyID := update.Message.MessageID
	if update.Message.Chat.IsPrivate() {
		replyID = 0
	}

	if message.File != nil {
		inst.sendFile(chatID, message, replyID)
		return
	}

	if message.PhotoURL != "" {
		inst.sendPhotoURL(chatID, message, replyID)
		return
	}

	inst.sendText(chatID, message, replyID)
}
