package commands

import (
	"fmt"
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
	"strconv"
	"strings"
)

// FeedbacksList command
func FeedbacksList(rep repository.IRepository) bot.Handler {
	return func(user *bot.User, chat *bot.Chat, msg *bot.Message) *bot.Message {
		countryCode, highwayID := parseHighwayID(msg.Text)

		list, err := rep.GetFeedbacksList(countryCode, highwayID)
		ulen := uint64(len(list))
		if err != nil || ulen == 0 {
			return nil
		}

		res := formatFeedback(&list[0])

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatFeedbackReplyMarkup(ulen, 1, countryCode, highwayID),
		}
	}
}

// FeedbackAnswer command
func FeedbackAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(data string) *bot.Message {
		arr := strings.Split(data, "|")
		countryCode := arr[0]
		highwayID := arr[1]

		index, err := strconv.ParseUint(arr[2], 10, 32)
		if err != nil {
			return nil
		}

		list, err := rep.GetFeedbacksList(countryCode, highwayID)
		length := uint64(len(list))
		if err != nil || length <= index {
			return nil
		}

		res := formatFeedback(&list[index])
		index = min(length, index+1)

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatFeedbackReplyMarkup(length, index, countryCode, highwayID),
		}
	}
}

func formatFeedbackReplyMarkup(length uint64, nextIndex uint64, countryCode string, highwayID string) *bot.InlineKeyboardMarkup {
	if length <= nextIndex {
		return nil
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			{
				{
					Text:         "Следующий отзыв",
					CallbackData: "/feedback:" + countryCode + "|" + highwayID + "|" + fmt.Sprint(nextIndex),
				},
			},
		},
	}
}

func parseHighwayID(text string) (string, string) {
	text = strings.ToLower(strings.ReplaceAll(strings.TrimSpace(text), "/", ""))

	// ru -> en
	text = strings.ReplaceAll(text, "а", "a")
	text = strings.ReplaceAll(text, "р", "p")
	text = strings.ReplaceAll(text, "м", "m")
	text = strings.ReplaceAll(text, "н", "h")

	countryCode := text[0:2]
	highwayID := text[2:3] + "-" + text[3:]

	return countryCode, highwayID
}

func formatFeedback(feedback *repository.Feedback) string {
	res := ""
	if feedback.Rating != 0 {
		res += fmt.Sprintf("*Ретинг:* %.1f\n", feedback.Rating)
	}
	if feedback.Date != "" {
		res += fmt.Sprintf("*Дата:* %s\n", feedback.Date)
	}
	if feedback.RoadPart != "" {
		res += fmt.Sprintf("*Участок:* %s\n", feedback.RoadPart)
	}
	res += "\n" + feedback.Text

	return res
}
