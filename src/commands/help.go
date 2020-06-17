package commands

import (
	"roadstatebot/src/bot"
	"roadstatebot/src/config"
)

// Help command
func Help(cfg *config.AppConfig) bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		return &bot.Message{Text: cfg.HelpText}
	}
}
