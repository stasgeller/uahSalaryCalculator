package service

import (
	"context"
	"uahSalaryBot/external/telegram"
	"uahSalaryBot/infrastructure/domain"
	"uahSalaryBot/usecase"

	"github.com/sirupsen/logrus"
)

//CommandsInfrastructure - basic command interface
type CommandsInfrastructure interface {
	StartCommand
}

//Handler - command handler
type Handler func(ctx context.Context, message *domain.Message) error

//Manager - manager for resolving commands with handlers.
type Manager struct {
	bot      usecase.TgBot
	commands map[string]Handler
}

//NewManagerServer - returns Command Manager
func NewManagerServer(bot usecase.TgBot, commands CommandsInfrastructure) *Manager {
	m := &Manager{bot: bot}

	m.commands = make(map[string]Handler)
	m.commands["start"] = commands.StartAction

	return m
}

//Run - receive messages from bot and run particular handler for each command
func (m *Manager) Run(ctx context.Context) {
	message := make(chan *domain.Message)

	go func() {
		for {
			select {
			case mes := <-message:
				if command, ok := telegram.CommandMap[mes.Command]; ok {
					if handler, ok := m.commands[command]; ok {
						err := handler(ctx, mes)
						logrus.Errorf("[%s] error: %s", mes.Command, err.Error())
						continue
					}
					continue
				}

				logrus.Errorf("[manager] error command: %s doesn't exist", mes.Command)
			case <-ctx.Done():
				m.bot.ShutDown()

				return
			}
		}
	}()

	m.bot.Listen(ctx, message)
}
