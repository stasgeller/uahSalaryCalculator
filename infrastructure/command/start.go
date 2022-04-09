package command

import (
	"context"
	"uahSalaryBot/infrastructure/domain"
)

type CurrentChat string

const ChatId CurrentChat = "currentChatId"

type HandlerActions interface {
	Use(context.Context, interface{}) error
}

type Start struct {
	usecase HandlerActions
}

func NewStart(uc HandlerActions) *Start {
	return &Start{uc}
}

//StartAction - returns Hello message with the list of available commands.
func (s *Start) StartAction(ctx context.Context, message *domain.Message) error {
	ctx = context.WithValue(ctx, ChatId, message.ChatID)

	if err := s.usecase.Use(ctx, message.User); err != nil {
		return err
	}

	return nil
}
