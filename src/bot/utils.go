package bot

import (
	"fmt"
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (inst *botStruct) toUser(from *tgbotapi.User) *User {
	return &User{ID: from.ID,
		FirstName:    from.FirstName,
		LastName:     from.LastName,
		UserName:     from.UserName,
		LanguageCode: from.LanguageCode,
		IsBot:        from.IsBot,
	}
}

func (inst *botStruct) toChat(from *tgbotapi.Chat) *Chat {
	photo := &ChatPhoto{}

	if from.Photo != nil {
		photo.BigFileID = from.Photo.BigFileID
		photo.SmallFileID = from.Photo.SmallFileID
	}

	return &Chat{
		ID:                  from.ID,
		Type:                from.Type,
		Title:               from.Title,
		UserName:            from.UserName,
		FirstName:           from.FirstName,
		LastName:            from.LastName,
		AllMembersAreAdmins: from.AllMembersAreAdmins,
		Description:         from.Description,
		InviteLink:          from.InviteLink,
		Photo:               photo,
	}
}

func (inst *botStruct) toMessage(from *tgbotapi.Message) *Message {
	return &Message{
		MessageID: from.MessageID,
		Date:      from.Date,
		Text:      from.Text,
	}
}

func (inst *botStruct) toTgFile(from *FileReader) *tgbotapi.FileReader {
	return &tgbotapi.FileReader{
		Name:   from.Name,
		Reader: from.Reader,
		Size:   from.Size,
	}
}

func (inst *botStruct) toTgInlineKeyboardMarkup(from *InlineKeyboardMarkup) *tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, row := range from.InlineKeyboard {
		var tgBtns []tgbotapi.InlineKeyboardButton
		for _, btn := range row {
			data := fmt.Sprint(btn.CallbackData)
			tgBtns = append(tgBtns, tgbotapi.InlineKeyboardButton{
				Text:         btn.Text,
				CallbackData: &data,
			})
		}

		rows = append(rows, tgBtns)
	}

	return &tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
}

func (inst *botStruct) findHandler(text string) Handler {
	for key, cb := range inst.commands {
		switch kv := key.(type) {
		case string:
			if key == text {
				return cb
			}
		case *regexp.Regexp:
			if kv.MatchString(text) {
				return cb
			}
		}
	}

	return inst.freeTextHandler
}

func (inst *botStruct) findAnswerHandler(command string) AnswerHandler {
	for key, cb := range inst.answers {
		if key == command {
			return cb
		}
	}

	return nil
}
