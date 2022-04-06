package domain

import tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Domain interface {
	Fill(upd *tgbot.Update)
}

type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	FirstName string
	LastName  string
}

func NewUser() *User {
	return &User{}
}

//Fill - convert telegram entity to domain entity
func (u *User) Fill(upd *tgbot.Update) {
	u.FirstName = upd.Message.Chat.FirstName
	u.LastName = upd.Message.Chat.LastName
	u.Username = upd.Message.Chat.UserName
}
