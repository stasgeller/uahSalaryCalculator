package service

import (
	"context"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramBot struct {
	*tgbot.BotAPI
}

func NewTgBot(token string) (*TelegramBot, error) {
	bot, err := tgbot.NewBotAPI(token)

	return &TelegramBot{bot}, err
}

func (tg *TelegramBot) Send(_ context.Context, m tgbot.Chattable) error {
	_, err := tg.BotAPI.Send(m)
	return err
}
