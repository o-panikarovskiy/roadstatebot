package commands

import (
	"roadstatebot/src/bot"
	"roadstatebot/src/config"
)

// Start command
func Start(cfg *config.AppConfig) bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		return &bot.Message{Text: cfg.StartText + "\n\n" + cfg.HelpText}
	}
}
