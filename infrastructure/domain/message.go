package domain

import (
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Message struct {
	MessageID int
	User      *User
	Date      int
	ChatID    int64
	Command   string
}

func NewMessage() *Message {
	return &Message{}
}

//Fill - convert telegram entity to domain entity
func (m *Message) Fill(u *tgbot.Update) {
	m.ChatID = u.Message.Chat.ID
	m.MessageID = u.Message.MessageID

	if u.Message.Command() != "" {
		m.Command = u.Message.Command()
	} else {
		m.Command = u.Message.Text
	}
}
