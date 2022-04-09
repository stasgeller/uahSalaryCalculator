package usecase

import (
	"context"
	"errors"
	"testing"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/infrastructure/domain"
	mock_usecase "uahSalaryBot/usecase/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestStart_Use(t *testing.T) {
	ctx := context.WithValue(context.Background(), command.ChatId, int64(2))
	ctrl := gomock.NewController(t)
	repo := mock_usecase.NewMockUserBase(ctrl)
	bot := mock_usecase.NewMockTgBot(ctrl)

	tests := []struct {
		name       string
		userDomain interface{}
		mockFunc   func(interface{})
		wantErr    error
	}{
		{
			name: "valid",
			userDomain: &domain.User{
				ID:        1,
				Username:  "test",
				FirstName: "test_first",
				LastName:  "test_last",
			},
			mockFunc: func(user interface{}) {
				repo.EXPECT().FindOrCreate(ctx, user).Return(nil)
				bot.EXPECT().Send(ctx, gomock.Any()).Return(nil)
			},
			wantErr: nil,
		},
		{
			name:       "unresolved domain",
			userDomain: domain.Message{},
			mockFunc:   func(interface{}) {},
			wantErr:    errors.New("unresolved user domain"),
		},
		{
			name: "nether found nor create",
			userDomain: &domain.User{
				ID:        1,
				Username:  "test",
				FirstName: "test_first",
				LastName:  "test_last",
			},
			mockFunc: func(user interface{}) {
				repo.EXPECT().FindOrCreate(ctx, user).Return(errors.New("msg err 1"))
			},
			wantErr: errors.New("msg err 1"),
		},
		{
			name: "not sent",
			userDomain: &domain.User{
				ID:        1,
				Username:  "test",
				FirstName: "test_first",
				LastName:  "test_last",
			},
			mockFunc: func(user interface{}) {
				repo.EXPECT().FindOrCreate(ctx, user).Return(nil)
				bot.EXPECT().Send(ctx, gomock.Any()).Return(errors.New("msg err 2"))
			},
			wantErr: errors.New("msg err 2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFunc(tt.userDomain)

			s := &Start{
				repository: repo,
				bot:        bot,
			}

			err := s.Use(ctx, tt.userDomain)

			assert.Equal(t, tt.wantErr, err)
		})
	}
}
