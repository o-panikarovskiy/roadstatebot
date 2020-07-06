package server

import (
	"log"
	"roadstatebot/src/bot"
)

// Instance server
type Instance struct {
	bot bot.IBot
}

// NewInstance create new Instance
func NewInstance(apiKey string) *Instance {
	return createDevInstase(apiKey)
}

// Run instanse
func (inst *Instance) Run() {
	inst.bot.Run()
	log.Println("Telegram bot started")
}

// Stop instanse
func (inst *Instance) Stop() {
	log.Println("Telegram bot stopped")
}
