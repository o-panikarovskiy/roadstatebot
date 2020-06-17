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
					CallbackData: "/list:" + country.Code,
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

// CountiesAnswer command
func CountiesAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(countryCode string) *bot.Message {
		list, err := rep.GetHighWayTypesList(countryCode)

		if err != nil {
			return &bot.Message{Text: err.Error()}
		}

		var rows [][]bot.InlineKeyboardButton
		for _, hwType := range list {
			row := []bot.InlineKeyboardButton{
				{
					Text:         hwType.Name,
					CallbackData: "/highway:" + countryCode + "|" + hwType.ID + "|" + hwType.Name,
				},
			}

			rows = append(rows, row)
		}

		return &bot.Message{
			Text: "Выберите тип трассы:",
			ReplyMarkup: &bot.InlineKeyboardMarkup{
				InlineKeyboard: rows,
			},
		}
	}
}
