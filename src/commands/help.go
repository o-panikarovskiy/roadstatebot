package commands

import (
	"roadstatebot/src/bot"
)

const helpText = `"Чтобы воспользоваться моей помощью введите одну из следующих команд:\n
/list - Показать список автодорог\n\nВы также можете ввести название трассы чтобы узнать отзывы о ней.
Например ввод /ruM4 покажет отзывы о трассе Москва – Ростов-на-Дону – Новороссийск"`

// Help command
func Help() bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		return &bot.Message{Text: helpText}
	}
}
