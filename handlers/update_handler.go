package handlers

import (
	"log"
	"telebot/models"
	"telebot/utils"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		panic(err)
	}
	return updates
}

func HandleUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message != nil { // Обработка обычных сообщений
		userState := models.GetUserState(int64(update.Message.From.ID))

		switch userState.State {
		case "root":
			userState.State = HandleCommand(bot, &update)
		case "year":
			HandleMessage(bot, &update)

		default:
			// Обработка других состояний, если необходимо
		}
	} else if update.CallbackQuery != nil { // Обработка нажатий на кнопки
		handleCallback(bot, update)
	}
}

func HandleCommand(bot *tgbotapi.BotAPI, update *tgbotapi.Update) string {
	message := update.Message
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Выберите год выпуска серии Hot Wheels \n ⬇️⬇️⬇️.")
			inlineKeyboard := utils.CreateKeyboardYear()
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)
			return "year" // Обновляем состояние
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не распознана.")
			bot.Send(msg)
		}
	}
	return "root"
}

func handleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	callback := update.CallbackQuery
	log.Printf("Кнопка нажата: %s", callback.Data)

	// Сохраняем ID предыдущего сообщения
	msgID := callback.Message.MessageID
	chatID := callback.Message.Chat.ID

	// Реакция на нажатие кнопок
	var responseMsg string
	switch callback.Data {
	case "th":
		responseMsg = "Вы выбрали Treasure Hunt!"
	case "supers":
		responseMsg = "Вы выбрали Supers!"
	case "2023":
		responseMsg = "Вы выбрали 2023!"
	case "2024":
		responseMsg = "Вы выбрали 2024!"
	default:
		responseMsg = "Неизвестный выбор."
	}

	// Отправка нового сообщения
	msg := tgbotapi.NewMessage(chatID, responseMsg)
	bot.Send(msg)

	// Удаление предыдущего сообщения
	bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: msgID,
	})

	// Подтверждение нажатия кнопки
	bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Вы выбрали: "+callback.Data))
}
