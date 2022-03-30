package command

import (
	"context"
	"fmt"
	"os"
	"uahSalaryBot/service"

	"github.com/enescakir/emoji"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgBot interface {
	Send(context.Context, tgbot.Chattable) error
}

//Start - returns Hello message with the list of available commands.
func Start(ctx context.Context, update tgbot.Update, bot *service.TelegramBot) {
	welcomeLetter := fmt.Sprintf(`Привет%s Добро пожаловать в %s`, emoji.SmilingFaceWithHalo, os.Getenv("BOT_NAME"))

	if err := bot.Send(ctx, tgbot.NewMessage(update.Message.Chat.ID, welcomeLetter)); err != nil {
		logrus.Errorf("[command]: could not send message - %s", err.Error())
	}
}
