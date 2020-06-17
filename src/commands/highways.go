package commands

import (
	"fmt"
	"roadstatebot/src/bot"
	"roadstatebot/src/repository"
	"strings"
)

// HighwayTypeAnswer command
func HighwayTypeAnswer(rep repository.IRepository) bot.AnswerHandler {
	return func(answer string) *bot.Message {
		answerData := strings.Split(answer, "|")
		countryCode := answerData[0]
		highwayTypeID := answerData[1]
		highwayName := answerData[2]

		list, err := rep.GetHighwaysList(countryCode, highwayTypeID)

		if err != nil {
			return &bot.Message{Text: err.Error()}
		}

		res := highwayName + ":\n"
		for _, highway := range list[:100] {
			if strings.HasPrefix(highway.ID, highwayTypeID) {
				res += fmt.Sprintf("/%s%s\t%s\t%.1f\n", countryCode, strings.Replace(highway.ID, "-", "", 1), highway.Name, highway.Rating)
			}
		}

		return &bot.Message{Text: res}
	}
}
