package autostrada

import (
	"fmt"
	"net/http"
	"roadstatebot/src/repository"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (s *autostrada) GetHighwaysList(countryCode string, highwayTypeID string) ([]repository.HighWay, error) {
	res, err := http.Get(domain + "/" + strings.ToLower(countryCode) + "/highways")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Невозможно получить список")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Невозможно получить список")
	}

	result := make([]repository.HighWay, 0)
	doc.Find(".highway-item").Each(func(i int, s *goquery.Selection) {
		highway := repository.HighWay{}

		highway.ID = parseHighwayID(s.Find("a.label-code").Attr("href"))
		highway.Name = s.Find("span.hidden-xs a.highwayLabel").Text()
		highway.Rating = parseHighwayRating(s.Find("span.rateNumber").Text())

		if strings.HasPrefix(highway.ID, highwayTypeID) {
			result = append(result, highway)
		}
	})

	return result, nil
}
