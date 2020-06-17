package commands

import (
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
)

// CountriesList command
func CountriesList(rep repository.IRepository) bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		list, err := rep.GetCountiesList()

		if err != nil {
			return &bot.Message{Text: err.Error()}
		}

		var rows [][]bot.InlineKeyboardButton
		for _, country := range list {
			row := []bot.InlineKeyboardButton{
				{
					Text:         country.Name,
					CallbackData: "/highways:" + country.Code + "|0|50",
				},
			}

			rows = append(rows, row)
		}

		return &bot.Message{
			Text: "Выберите страну:",
			ReplyMarkup: &bot.InlineKeyboardMarkup{
				InlineKeyboard: rows,
			},
		}
	}
}
