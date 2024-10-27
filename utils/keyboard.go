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
			tgbotapi.NewInlineKeyboardButtonData("1", "1"),
			tgbotapi.NewInlineKeyboardButtonData("2", "2"),
			tgbotapi.NewInlineKeyboardButtonData("3", "3"),
			tgbotapi.NewInlineKeyboardButtonData("4", "4"),
			tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("6", "6"),
			tgbotapi.NewInlineKeyboardButtonData("7", "7"),
			tgbotapi.NewInlineKeyboardButtonData("8", "8"),
			tgbotapi.NewInlineKeyboardButtonData("9", "9"),
			tgbotapi.NewInlineKeyboardButtonData("10", "10"),
		), tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("11", "11"),
			tgbotapi.NewInlineKeyboardButtonData("12", "12"),
			tgbotapi.NewInlineKeyboardButtonData("13", "13"),
			tgbotapi.NewInlineKeyboardButtonData("14", "14"),
			tgbotapi.NewInlineKeyboardButtonData("15", "15"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Все фото", "all"),
		),
	)
}

func CreateKeyboard() tgbotapi.ReplyKeyboardMarkup {
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("2023-TH"),
			tgbotapi.NewKeyboardButton("2023-Supers"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("2024-TH"),
			tgbotapi.NewKeyboardButton("2024-Supers"),
		),
	)

	// Устанавливаем свойства клавиатуры
	keyboard.ResizeKeyboard = true
	keyboard.OneTimeKeyboard = false // Делаем клавиатуру постоянной

	return keyboard
}
