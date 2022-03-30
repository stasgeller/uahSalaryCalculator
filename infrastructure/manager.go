package infrastructure

import (
	"context"
	"os"
	"uahSalaryBot/infrastructure/command"
	"uahSalaryBot/service"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler func(ctx context.Context, update tgbot.Update, bot *service.TelegramBot)

type Manager struct {
	bot      *command.TgBot
	commands map[string]Handler
}

//NewManager - returns Command Manager
func NewManager(bot command.TgBot) *Manager {
	m := &Manager{bot: &bot}

	m.commands = make(map[string]Handler)
	m.commands["/start"] = command.Start

	return m
}

//Run - receive messages from bot and run particular handler for each command
func (m *Manager) Run(ctx context.Context, bot *service.TelegramBot) {
	go func() {
		<-ctx.Done()
		bot.StopReceivingUpdates()
		os.Exit(1)
	}()

	ucfg := tgbot.NewUpdate(0)
	ucfg.Timeout = 30

	updates := bot.GetUpdatesChan(ucfg)
	for update := range updates {
		if update.Message != nil {
			if handler, ok := m.command(update.Message.Command()); !ok {
				continue
			} else {
				handler(ctx, update, bot)
			}
		}
	}
}

//command - returns Handler which can be called for certain CommandName
//if second parameter equal false then Handler doesn't exist.
func (m *Manager) command(name string) (Handler, bool) {
	if name == "" {
		return nil, false
	}

	if c, ok := m.commands[name]; !ok {
		return nil, false
	} else {
		return c, true
	}
}
