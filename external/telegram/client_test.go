package telegram

import (
	"context"
	"errors"
	"testing"
	mock_telegram "uahSalaryBot/external/telegram/mock"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestTelegramBot_Send(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_telegram.NewMockTgClient(ctrl)

	ctx := context.Background()
	msg := tgbot.NewMessage(1, "some text")

	tests := []struct {
		name    string
		wantErr error
	}{
		{
			name:    "sent",
			wantErr: nil,
		},
		{
			name:    "fail",
			wantErr: errors.New("err msg"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg.ReplyMarkup = Keyboard()
			m.EXPECT().Send(msg).Return(tgbot.Message{}, tt.wantErr)

			tg := &TelegramBot{m}

			err := tg.Send(ctx, msg)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
