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

func CreateKeyboardSeries() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("TH", "th"),
			tgbotapi.NewInlineKeyboardButtonData("Supers", "supers"),
		),
	)
}

func CreateKeyboardPhotos() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Назад", "back"),
			tgbotapi.NewInlineKeyboardButtonData("Все фото", "all"),
		),
	)
}
