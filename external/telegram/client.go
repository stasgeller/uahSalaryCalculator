package telegram

import (
	"context"
	"uahSalaryBot/infrastructure/domain"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//TelegramBot - custom telegram client
type TelegramBot struct {
	client *tgbot.BotAPI
}

//NewTgBot - initialize telegram bot client.
func NewTgBot(token string) (*TelegramBot, error) {
	bot, err := tgbot.NewBotAPI(token)

	return &TelegramBot{client: bot}, err
}

//Send - sends messages with keyboard
func (tg *TelegramBot) Send(_ context.Context, m tgbot.MessageConfig) error {
	m.ReplyMarkup = Keyboard()

	_, err := tg.client.Send(m)
	return err
}

//ShutDown - stop receiving messages from telegram channel.
func (tg *TelegramBot) ShutDown() {
	tg.client.StopReceivingUpdates()
}

//Listen - listening telegram channel
func (tg *TelegramBot) Listen(_ context.Context, chat chan *domain.Message) {
	ucfg := tgbot.NewUpdate(0)
	ucfg.Timeout = 30

	for update := range tg.client.GetUpdatesChan(ucfg) {
		if update.Message != nil {
			message := domain.NewMessage()
			message.Fill(&update)

			message.User = domain.NewUser()
			message.User.Fill(&update)

			chat <- message
		} else {
			continue
		}
	}
}
