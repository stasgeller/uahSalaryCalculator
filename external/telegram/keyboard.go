package telegram

import (
	"fmt"

	"github.com/enescakir/emoji"
	tgbot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

//CommandMap uses to set up bot keyboard and map it commands. Key is button text, value is command string
var CommandMap = map[string]string{
	//do not touch init command
	"start":              StartButtonCommand,
	StartButtonKey:       StartButtonCommand,
	BankAuthorizationKey: BankAuthorizationCommand,
}

var StartButtonKey = fmt.Sprintf("%sО боте", emoji.SmilingFace)
var StartButtonCommand = "start"

var BankAuthorizationKey = fmt.Sprintf("%sПодключить банк", emoji.Bank)
var BankAuthorizationCommand = "monobank_auth"

type keyboard tgbot.ReplyKeyboardMarkup

//Keyboard - returns keyboard for customer
func Keyboard() *keyboard {
	buttons := tgbot.NewReplyKeyboard(
		tgbot.NewKeyboardButtonRow(
			tgbot.NewKeyboardButton(StartButtonKey),
		),
		tgbot.NewKeyboardButtonRow(
			tgbot.NewKeyboardButton(BankAuthorizationKey),
		),
	)

	var k keyboard
	k = keyboard(buttons)

	return &k
}
