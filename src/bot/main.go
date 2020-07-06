package bot

import (
	"regexp"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type botStruct struct {
	api             *tgbotapi.BotAPI
	commands        map[interface{}]Handler
	answers         map[string]AnswerHandler
	freeTextHandler Handler
}

// New create new instanse of IBot
func New(apiKey string) (IBot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)

	inst := &botStruct{
		api:      bot,
		commands: make(map[interface{}]Handler),
		answers:  make(map[string]AnswerHandler),
	}

	return inst, err
}

// BotName return bot-self name
func (inst *botStruct) BotName() string {
	return inst.api.Self.UserName
}

func (inst *botStruct) On(cmd string, handler Handler) {
	inst.commands[cmd] = handler
}

func (inst *botStruct) OnRegexp(reg *regexp.Regexp, handler Handler) {
	inst.commands[reg] = handler
}

func (inst *botStruct) OnText(handler Handler) {
	inst.freeTextHandler = handler
}

func (inst *botStruct) OnAnswer(cmd string, handler AnswerHandler) {
	inst.answers[cmd] = handler
}
