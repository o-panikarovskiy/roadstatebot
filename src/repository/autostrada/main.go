package autostrada

import (
	"roadstatebot/src/repository"
)

const domain = "https://autostrada.info"

type autostrada struct {
}

// New func
func New() repository.IRepository {
	return &autostrada{}
}
