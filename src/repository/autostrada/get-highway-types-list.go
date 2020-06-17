package autostrada

import (
	"roadstatebot/src/repository"
	"strings"
)

var dic = map[string][]repository.HighWayType{
	"ru": {
		{
			ID:   "M",
			Name: "ФЕДЕРАЛЬНЫЕ",
		},
		{
			ID:   "P",
			Name: "РЕГИОНАЛЬНЫЕ",
		},
		{
			ID:   "A",
			Name: "СОЕДИНИТЕЛЬНЫЕ",
		},
	},
	"ua": {
		{
			ID:   "M",
			Name: "МЕЖДУНАРОДНЫЕ",
		},
		{
			ID:   "H",
			Name: "НАЦИОНАЛЬНЫЕ",
		},
		{
			ID:   "P",
			Name: "РЕГИОНАЛЬНЫЕ",
		},
	},
	"by": {
		{
			ID:   "M",
			Name: "МАГИСТРАЛИ",
		},
	},
}

func (s *autostrada) GetHighWayTypesList(countryCode string) ([]repository.HighWayType, error) {
	return dic[strings.ToLower(countryCode)], nil
}
