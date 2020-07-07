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
	Emotion     string `json:"emotion"`
	URL         string `json:"url"`
	Action      string `json:"action"`
}

const anfisaHelp = "Для начала диалога используй восклицательный знак.\nНапример так:\n! привет!"

var anfisaError = []string{"🙄", "😴", "🙄", "😊"}

// AnfisaChat func
func AnfisaChat(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
	text := strings.TrimSpace(trimFirstRune(msg.Text))
	if text == "" {
		return &bot.Message{Text: anfisaHelp}
	}

	req := url.Values{
		"query": {fmt.Sprintf(`{"ask": "%s","key":"","userid":"%v"}`, text, chat.ID)},
	}

	resp, err := http.PostForm("https://aiproject.ru/api/", req)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("anfisa http error: %v - %v", resp.StatusCode, err)
		return &bot.Message{Text: getRandValueInArr(anfisaError)}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &bot.Message{}
	}

	answer := anfisaAnswer{}
	err = json.Unmarshal(body, &answer)
	if err != nil {
		return &bot.Message{}
	}

	if answer.Status != 1 {
		log.Printf("anfisa status error: %s", string(body))
		return &bot.Message{Text: getRandValueInArr(anfisaError)}
	}

	return &bot.Message{
		Text:     answer.Aiml,
		PhotoURL: answer.URL,
	}
}
