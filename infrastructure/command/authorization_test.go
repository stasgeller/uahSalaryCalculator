package command

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	mock_command "uahSalaryBot/infrastructure/command/mock"
	"uahSalaryBot/infrastructure/domain"
)

func TestAuth_Authorization(t *testing.T) {
	ctrl := gomock.NewController(t)
	m := mock_command.NewMockHandlerActions(ctrl)
	ctx := context.Background()

	tests := []struct {
		name    string
		message *domain.Message
		wantErr error
	}{
		{
			name: "valid",
			message: &domain.Message{
				ChatID: 1,
			},
			wantErr: nil,
		},
		{
			name: "error",
			message: &domain.Message{
				ChatID: 1,
			},
			wantErr: errors.New("err msg"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m.EXPECT().Use(context.WithValue(ctx, ChatId, tt.message.ChatID), tt.message.User).Return(tt.wantErr)

			s := &Auth{m}
			err := s.Authorization(ctx, tt.message)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
