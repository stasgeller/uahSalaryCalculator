package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/infrastructure/domain"
	mock_usecase "uahSalaryBot/usecase/mock"
)

func TestAuth_Use(t *testing.T) {
	ctx := context.WithValue(context.Background(), command.ChatId, int64(2))
	ctrl := gomock.NewController(t)
	botMock := mock_usecase.NewMockTgBot(ctrl)

	var m interface{}
	m = &domain.Message{}
	notApplicable, _ := m.(*domain.User)

	tests := []struct {
		name       string
		userDomain interface{}
		mockFunc   func(error)
		wantErr    error
	}{
		{
			name: "Valid",
			userDomain: &domain.User{
				ID:        1,
				Username:  "test",
				FirstName: "test_first",
				LastName:  "test_last",
			},
			mockFunc: func(err error) {
				botMock.EXPECT().Send(ctx, gomock.Any()).Return(err)
			},
			wantErr: nil,
		},
		{
			name:       "Not applicable domain",
			userDomain: &domain.Message{},
			mockFunc:   func(err error) {},
			wantErr:    &NotApplicableDomain{notApplicable},
		},
		{
			name: "Sending error",
			userDomain: &domain.User{
				ID:        1,
				Username:  "test",
				FirstName: "test_first",
				LastName:  "test_last",
			},
			mockFunc: func(err error) {
				botMock.EXPECT().Send(ctx, gomock.Any()).Return(err)
			},
			wantErr: errors.New("err msg"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.wantErr)
			a := &Auth{botMock}

			err := a.Use(ctx, tt.userDomain)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
