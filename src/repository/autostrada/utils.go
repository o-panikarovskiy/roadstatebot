package autostrada

import (
	"strconv"
	"strings"
)

func parseHighwayID(href string, ok bool) string {
	if !ok {
		return ""
	}

	paths := strings.Split(href, "/")
	return strings.ToUpper(paths[len(paths)-1])
}

func parseHighwayRating(text string) float64 {
	rtext := strings.ReplaceAll(text, ",", ".")
	rating, err := strconv.ParseFloat(rtext, 64)

	if err == nil {
		return rating
	}

	return 0.0
}
