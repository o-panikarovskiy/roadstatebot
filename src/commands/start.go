package commands

import (
	"roadstatebot/src/bot"
)

// Start command
func Start() bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		return &bot.Message{
			Text: `"Привет! Я бот который знает состояние дорог твоей страны.\n\n
			Я умею показывать оценки и отзывы автовладельцев о трассах."\n\n` + helpText}
	}
}
