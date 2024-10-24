package utils

import "github.com/go-telegram-bot-api/telegram-bot-api"

func CreateKeyboardYear() tgbotapi.ReplyKeyboardMarkup {
	return tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("2023"),
			tgbotapi.NewKeyboardButton("2024"),
		),
	)
}
