package autostrada

import (
	"fmt"
	"net/http"
	"roadstatebot/src/repository"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (s *autostrada) GetFeedbacksList(countryCode string, highwayID string) ([]repository.Feedback, error) {
	url := domain + "/" + strings.ToLower(countryCode) + "/highway/" + strings.ToUpper(highwayID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Невозможно получить отзывы")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить отзывы")
	}

	result := make([]repository.Feedback, 0)
	doc.Find(".b-reviews .item").Each(func(i int, s *goquery.Selection) {
		var feedback = repository.Feedback{}

		feedback.Rating = parseHighwayRating(s.Find(".b-rate.visible-xs .sharpen").Text())
		feedback.Date = strings.TrimSpace(s.Find(".reviewDate.hidden-xs").Text())
		feedback.RoadPart = strings.TrimSpace(s.Find("p.hidden-xs b").Text())
		feedback.Text = strings.TrimSpace(s.Find("p.comment").Text())

		if feedback.Text != "" {
			result = append(result, feedback)
		}
	})

	return result, nil
}
