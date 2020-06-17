package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// AppConfig main app config
type AppConfig struct {
	Env           string `json:"env"`
	TelegramToken string `json:"telegramToken"`
	StartText     string `json:"startText"`
	HelpText      string `json:"helpText"`
}

const (
	// DevMode indicates mode is debug.
	DevMode = "dev"
	// ProdMode indicates mode is production.
	ProdMode = "prod"
	// TestMode indicates mode is test.
	TestMode = "test"
)

// IsDev returns true if env in DevMode
func (c *AppConfig) IsDev() bool {
	return c.Env == DevMode
}

// IsProd returns true if env in ProdMode
func (c *AppConfig) IsProd() bool {
	return c.Env == ProdMode
}

// IsTest returns true if env in TestMode
func (c *AppConfig) IsTest() bool {
	return c.Env == TestMode
}

// NewDefaultConfig parses command line arguments and read json config
func NewDefaultConfig(path string) *AppConfig {
	return readConfigFile(path)
}

func readConfigFile(path string) *AppConfig {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Panicln(err)
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panicln(err)
	}

	var result AppConfig
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		log.Panicln(err)
	}

	return &result
}
