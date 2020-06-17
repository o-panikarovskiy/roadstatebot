package autostrada

import (
	"roadstatebot/src/repository"
)

func (s *autostrada) GetCountiesList() ([]repository.Country, error) {
	return []repository.Country{
		{
			Name: "Россия",
			Code: "ru",
		},
		{
			Name: "Украина",
			Code: "ua",
		},
		{
			Name: "Беларусь",
			Code: "by",
		},
	}, nil
}
