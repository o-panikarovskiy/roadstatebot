package commands

type userState struct {
	CallCounter     int
	PrevCommad      string
	PrevCommandArgs interface{}
}

var commadState = make(map[string]*userState)
