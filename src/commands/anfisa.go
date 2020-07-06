package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"roadstatebot/src/bot"
	"strings"
)

type anfisaAnswer struct {
	Status      int    `json:"status"`
	Aiml        string `json:"aiml"`
	Description string `json:"description"`
}

// AnfisaChat func
func AnfisaChat(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
	text := strings.TrimSpace(trimFirstRune(msg.Text))
	if text == "" {
		return &bot.Message{Text: "Для начала диалога используй восклицательный знак.\nНапример так:\n! привет!"}
	}

	req := url.Values{
		"query": {fmt.Sprintf(`{"ask": "%s","key":"","userid":"%v"}`, text, chat.ID)},
	}

	resp, err := http.PostForm("https://aiproject.ru/api/", req)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("anfisa response error: %v - %v \n", resp.StatusCode, err)
		return &bot.Message{Text: "Я не знаю что тут сказать..."}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("anfisa read body error:  %v \n", err)
		return &bot.Message{}
	}

	answer := anfisaAnswer{}
	err = json.Unmarshal(body, &answer)
	if err != nil {
		log.Printf("anfisa parse json error:  %v \n", err)
		return &bot.Message{}
	}

	if answer.Status != 1 {
		log.Printf("anfisa status error:  %s \n", answer.Description)
		return &bot.Message{}
	}

	return &bot.Message{Text: answer.Aiml}
}
