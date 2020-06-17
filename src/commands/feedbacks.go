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
		if err != nil || len(list) == 0 {
			return nil
		}

		res := formatFeedback(&list[0])

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatReplyMarkup(len(list), 1, countryCode, highwayID),
		}
	}
}

// FeedbackAnswer command
func FeedbackAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(data string) *bot.Message {
		arr := strings.Split(data, "|")
		countryCode := arr[0]
		highwayID := arr[1]
		indexStr := arr[2]
		index, err := strconv.ParseUint(indexStr, 10, 32)
		if err != nil {
			return nil
		}

		list, err := rep.GetFeedbacksList(countryCode, highwayID)
		ulen := uint64(len(list))
		if err != nil || ulen <= index {
			return nil
		}

		res := formatFeedback(&list[index])

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatReplyMarkup(len(list), int(index+1), countryCode, highwayID),
		}
	}
}

func formatReplyMarkup(maxLen int, nextIndex int, countryCode string, highwayID string) *bot.InlineKeyboardMarkup {
	if maxLen <= nextIndex {
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
		res += fmt.Sprintf("Ретинг: %.1f \n", feedback.Rating)
	}
	if feedback.Date != "" {
		res += fmt.Sprintf("Дата: %s \n", feedback.Date)
	}
	if feedback.RoadPart != "" {
		res += fmt.Sprintf("Участок: %s \n", feedback.RoadPart)
	}
	res += feedback.Text

	return res
}
