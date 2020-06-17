package server

import (
	"log"
	"roadstatebot/src/bot"
	"roadstatebot/src/config"
)

// Instance server
type Instance struct {
	bot bot.IBot
	cfg *config.AppConfig
}

// NewInstance create new Instance
func NewInstance(cfg *config.AppConfig) *Instance {
	if cfg.IsDev() {
		return createDevInstase(cfg)
	}
	return nil
}

// Run instanse
func (inst *Instance) Run() {
	inst.bot.Run()
	log.Print("Telegram bot started")
}

// Stop instanse
func (inst *Instance) Stop() {
	log.Print("Telegram bot stopped")
}
