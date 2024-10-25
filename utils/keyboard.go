package utils

import "github.com/go-telegram-bot-api/telegram-bot-api"

func CreateKeyboardYear() tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("2023", "2023"),
			tgbotapi.NewInlineKeyboardButtonData("2024", "2024"),
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
			tgbotapi.NewInlineKeyboardButtonData("Все фото", "all"),
		),
	)
}
