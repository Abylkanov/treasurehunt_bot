package handlers

import (
	"telebot/models"
	"telebot/utils"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleMessage(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	message := update.Message
	userState := models.GetUserState(int64(message.From.ID))

	switch message.Text {
	case "2023":
		Handle2023(bot, update)
		// Обновление состояния здесь по необходимости
	case "2024":
		msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите серию Treasure Hunt:")
		inlineKeyboard := utils.CreateKeyboardSeries()
		msg.ReplyMarkup = inlineKeyboard
		bot.Send(msg)
		// Обновление состояния
		userState.State = "year" // или другое состояние
	default:
		msg := tgbotapi.NewMessage(message.Chat.ID, "К доступны только 2023-2024 год")
		bot.Send(msg)
	}
}

func Handle2023(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	message := update.Message
	msg := tgbotapi.NewMessage(message.Chat.ID, "Выберите серию Treasure Hunt:")
	inlineKeyboard := utils.CreateKeyboardSeries()
	msg.ReplyMarkup = inlineKeyboard
	bot.Send(msg)

	if update.CallbackQuery != nil { // Если есть нажатие кнопки
		callback := update.CallbackQuery

		switch callback.Data {
		case "th":
			msg := tgbotapi.NewMessage(message.Chat.ID, list2023th)
			inlineKeyboard := utils.CreateKeyboardPhotos()
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)

		case "supers":
			msg := tgbotapi.NewMessage(message.Chat.ID, "not Ready")
			inlineKeyboard := utils.CreateKeyboardPhotos()
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "press the buton")
			bot.Send(msg)
		}

		// Удаление предыдущего сообщения (опционально)
		bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Вы выбрали: "+callback.Data))
	}
}
