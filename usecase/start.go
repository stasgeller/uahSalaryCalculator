package usecase

import (
	"context"
	"fmt"
	"os"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/infrastructure/domain"
	"uahSalaryBot/infrastructure/repository"

	"github.com/enescakir/emoji"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type TgBot interface {
	Send(context.Context, tgbot.MessageConfig) error
	ShutDown()
	Listen(context.Context, chan *domain.Message)
}

//UserBase - base user repository interface
type UserBase interface {
	Create(context.Context) error
	ReadOne(context.Context) error
	Update(context.Context) error
	Delete(context.Context) error
	ModelUser
	FindUser
}

//ModelUser - uses to handle model
type ModelUser interface {
	SetModel(*domain.User) *repository.User
}

//FindUser - finder for user repository
type FindUser interface {
	FindOrCreate(context.Context) error
}

type Repositories struct {
	User UserBase
}

type Clients struct {
	client       TgBot
	repositories *Repositories
}

func NewClients(c TgBot, r *Repositories) *Clients {
	return &Clients{
		client:       c,
		repositories: r,
	}
}

type Start struct {
	repository UserBase
	bot        TgBot
}

func StartCase(cs *Clients) *Start {
	return &Start{
		cs.repositories.User,
		cs.client,
	}
}

//Use - performs use case for certain Clients. Use services if it needs.
func (s *Start) Use(ctx context.Context, userDomain interface{}) error {
	chat := ctx.Value(command.ChatId)
	chatId := chat.(int64)

	user := userDomain.(*domain.User)

	if err := s.repository.SetModel(user).FindOrCreate(ctx); err != nil {
		return err
	}

	welcomeLetter := fmt.Sprintf(
		`Привет%s Добро пожаловать в %s. 
Пожалуйста выберите 1 из доступных действий в меню%s`,
		emoji.SmilingFaceWithHalo,
		os.Getenv("BOT_NAME"),
		emoji.SmilingFace,
	)

	if err := s.bot.Send(ctx, tgbot.NewMessage(chatId, welcomeLetter)); err != nil {
		logrus.Errorf("[command]: could not send message - %s", err.Error())
	}

	return nil
}
