package commands

import (
	"fmt"
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
	"strconv"
	"strings"
)

// HighwayListAnswer command
func HighwayListAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(answer string) *bot.Message {
		arr := strings.Split(answer, "|")
		countryCode := arr[0]
		highwayTypeID := arr[1]
		step := uint64(50)

		indexFrom, err := strconv.ParseUint(arr[2], 10, 32)
		if err != nil {
			return nil
		}

		indexTo, err := strconv.ParseUint(arr[3], 10, 32)
		if err != nil {
			return nil
		}

		list, err := rep.GetHighwaysList(countryCode, highwayTypeID)
		ulen := uint64(len(list))
		if err != nil || ulen == 0 {
			return nil
		}

		if indexTo >= ulen {
			indexTo = ulen - 1
		}

		res := ""
		for _, highway := range list[indexFrom:indexTo] {
			res += fmt.Sprintf("/%s%s\t%s\t%.1f\n", countryCode, strings.ToLower(strings.Replace(highway.ID, "-", "", 1)), highway.Name, highway.Rating)
		}

		return &bot.Message{
			Text:        res,
			ReplyMarkup: formatHighwayListReplyMarkup(ulen, indexFrom+step, indexTo+step, countryCode, highwayTypeID),
		}
	}
}

func formatHighwayListReplyMarkup(maxLen uint64, indexFrom uint64, indexTo uint64, countryCode string, highwayTypeID string) *bot.InlineKeyboardMarkup {
	if maxLen <= indexFrom || indexTo >= maxLen {
		return nil
	}

	return &bot.InlineKeyboardMarkup{
		InlineKeyboard: [][]bot.InlineKeyboardButton{
			{
				{
					Text:         "Показать еще",
					CallbackData: "/highways:" + countryCode + "|" + highwayTypeID + "|" + fmt.Sprint(indexFrom) + "|" + fmt.Sprint(indexTo),
				},
			},
		},
	}
}
