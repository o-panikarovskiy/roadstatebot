package commands

import (
	"fmt"
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
	"strconv"
	"strings"
)

const pageSize uint64 = 50

// HighwayListAnswer command
func HighwayListAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(answer string) *bot.Message {
		arr := strings.Split(answer, "|")
		countryCode := arr[0]

		indexFrom, err := strconv.ParseUint(arr[1], 10, 32)
		if err != nil {
			return nil
		}

		indexTo, err := strconv.ParseUint(arr[2], 10, 32)
		if err != nil {
			return nil
		}

		list, err := rep.GetHighwaysList(countryCode, "")
		length := uint64(len(list))
		if err != nil || length <= indexFrom {
			return nil
		}

		indexTo = min(indexTo, length)

		res := ""
		for _, highway := range list[indexFrom:indexTo] {
			res += fmt.Sprintf("/%s%s\t%s\t%.1f\n", countryCode, strings.ToLower(strings.Replace(highway.ID, "-", "", 1)), highway.Name, highway.Rating)
		}

		indexFrom = indexTo
		indexTo = min(length, indexTo+pageSize)

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatHighwayListReplyMarkup(length, indexFrom, indexTo, countryCode),
		}
	}
}

func formatHighwayListReplyMarkup(length uint64, indexFrom uint64, indexTo uint64, countryCode string) *bot.InlineKeyboardMarkup {
	if length <= indexFrom {
		return nil
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			{
				{
					Text:         "Показать еще",
					CallbackData: "/highways:" + countryCode + "|" + fmt.Sprint(indexFrom) + "|" + fmt.Sprint(indexTo),
				},
			},
		},
	}
}
