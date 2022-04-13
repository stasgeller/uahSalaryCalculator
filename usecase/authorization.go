package usecase

import (
	"context"
	"fmt"
	"github.com/enescakir/emoji"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/infrastructure/domain"
)

const monobankAuthUrl = "https://api.monobank.ua/"

type Auth struct {
	bot TgBot
}

func AuthorizationUseCase(c *Clients) *Auth {
	return &Auth{
		bot: c.bot,
	}
}

//Use - sends message with available banks for authorization
func (a *Auth) Use(ctx context.Context, userDomain interface{}) error {
	chat := ctx.Value(command.ChatId)
	chatId := chat.(int64)

	user, ok := userDomain.(*domain.User)
	if !ok {
		return &NotApplicableDomain{user}
	}

	msg := tgbot.NewMessage(chatId, fmt.Sprintf(`
Для авторизации пожалуйста выберите 1 из доступных банков ниже. Следуйте инструкциям:

%s*Монобанк* - для подключения Монобанка перейдите по ссылке ниже и авторизуйтесь.
Получите ключ доступа и скопируйте его сюда.`, emoji.ExclamationMark))
	msg.ReplyMarkup = authLinks()
	if err := a.bot.Send(ctx, msg); err != nil {
		return err
	}

	return nil
}

//authLinks - returns button list for bank connection
func authLinks() tgbot.InlineKeyboardMarkup {
	return tgbot.NewInlineKeyboardMarkup(
		tgbot.NewInlineKeyboardRow(
			tgbot.NewInlineKeyboardButtonURL("Монобанк", monobankAuthUrl),
		),
	)
}
