package bot

import "regexp"

// User is a user on Telegram.
type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`     // optional
	UserName     string `json:"username"`      // optional
	LanguageCode string `json:"language_code"` // optional
	IsBot        bool   `json:"is_bot"`        // optional
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID                  int64      `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`                          // optional
	UserName            string     `json:"username"`                       // optional
	FirstName           string     `json:"first_name"`                     // optional
	LastName            string     `json:"last_name"`                      // optional
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"` // optional
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"` // optional
	InviteLink          string     `json:"invite_link,omitempty"` // optional
}

// Message struct
type Message struct {
	MessageID   int    `json:"message_id"`
	Date        int    `json:"date"`
	Text        string `json:"text"`
	ReplyMarkup *InlineKeyboardMarkup
}

// InlineKeyboardMarkup is a custom keyboard presented for an inline bot.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton is a button within a custom keyboard for inline query responses.
type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data,omitempty"` // optional
}

// Handler  function
type Handler func(user *User, chat *Chat, msg *Message) *Message

// AnswerHandler callback function
type AnswerHandler func(answer string) *Message

// IBot interface
type IBot interface {
	// BotName returns bot username
	BotName() string

	// Run start listen
	Run()

	// On add string command listener
	On(cmd string, handler Handler)

	// OnRegexp add regexp command listener
	OnRegexp(cmd *regexp.Regexp, handler Handler)

	// OnText add free text listener
	OnText(handler Handler)

	// OnAnswer add answer listener
	OnAnswer(cmd string, handler AnswerHandler)
}
