package service

import (
	"context"
	"uahSalaryBot/external/db"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/infrastructure/domain"
	"uahSalaryBot/infrastructure/repository"
	"uahSalaryBot/usecase"
)

//NewRepositories - initialize all repositories
func NewRepositories(c *db.DbClient) *usecase.Repositories {
	return &usecase.Repositories{
		User: repository.NewUserRepository(c),
	}
}

//UseCases - stores all use cases for each command
type UseCases struct {
	Start command.HandlerActions
}

//NewUseCases - initialize all use cases
func NewUseCases(cs *usecase.Clients) *UseCases {
	return &UseCases{
		Start: usecase.StartCase(cs),
	}
}

//StartCommand - stores methods for Start action
type StartCommand interface {
	StartAction(context.Context, *domain.Message) error
}

//Commands - stores all available commands
type Commands struct {
	StartCommand
}

//NewCommands - initialize all commands.
func NewCommands(uc *UseCases) *Commands {
	return &Commands{
		StartCommand: command.NewStart(uc.Start),
	}
}
