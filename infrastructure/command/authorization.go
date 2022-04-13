package command

import (
	"context"
	"uahSalaryBot/infrastructure/domain"
)

type Auth struct {
	usecase HandlerActions
}

func NewAuth(uc HandlerActions) *Auth {
	return &Auth{usecase: uc}
}

//Authorization - performs command to get all available banks for this bot
func (a *Auth) Authorization(ctx context.Context, message *domain.Message) error {
	ctx = context.WithValue(ctx, ChatId, message.ChatID)

	if err := a.usecase.Use(ctx, message.User); err != nil {
		return err
	}

	return nil
}
