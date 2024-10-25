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

		if update.Message.IsCommand() {
			HandleCommand(bot, update.Message, userState)
		}
	} else if update.CallbackQuery != nil { // Обработка нажатий на кнопки
		handleCallback(bot, update.CallbackQuery)
	}
}

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, userState *models.UserState) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Выберите год выпуска серии Hot Wheels \n ⬇️⬇️⬇️.")
			inlineKeyboard := utils.CreateKeyboardYear()
			msg.ReplyMarkup = inlineKeyboard
			bot.Send(msg)
			// Инициализация состояния пользователя
			userState.Data = make(map[string]interface{})
			userState.Data["selected_year"] = nil
			userState.Data["selected_series"] = nil
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не распознана.")
			bot.Send(msg)
		}
	}
}

func handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	log.Printf("Кнопка нажата: %s", callback.Data)

	// Сохраняем ID предыдущего сообщения
	msgID := callback.Message.MessageID
	chatID := callback.Message.Chat.ID

	responseMsg := "Неизвестный выбор."
	userState := models.GetUserState(int64(callback.From.ID))

	switch {
	case callback.Data == "2023" || callback.Data == "2024":
		userState.Data["selected_year"] = callback.Data
		responseMsg = "Вы выбрали " + callback.Data + "! Теперь выберите тип."
		msg := tgbotapi.NewMessage(chatID, responseMsg)
		msg.ReplyMarkup = utils.CreateKeyboardSeries()
		bot.Send(msg)

	case callback.Data == "th" || callback.Data == "supers":
		if year, ok := userState.Data["selected_year"].(string); ok && year != "" {
			userState.Data["selected_series"] = callback.Data
			responseMsg = getSelectedMessage(year, callback.Data)
			msg := tgbotapi.NewMessage(chatID, responseMsg)
			msg.ReplyMarkup = utils.CreateKeyboardPhotos()
			bot.Send(msg)
		} else {
			responseMsg = "Сначала выберите год."
			msg := tgbotapi.NewMessage(chatID, responseMsg)
			bot.Send(msg)

		}
	case callback.Data == "all":
		if year, ok := userState.Data["selected_year"].(string); ok && year != "" {
			if series, ok := userState.Data["selected_series"].(string); ok && series != "" {
				files := getPhoto(year, series, 0) // Получаем все файлы

				// Проверяем, что files не nil и это мапа
				if files != nil {
					for key, value := range files.(map[string]string) { // Приводим к нужному типу
						err := SendPhoto(bot, chatID, value, key)
						if err != nil {
							log.Printf("Ошибка при отправке фото %s: %s", value, err)
						}
					}
				}
			}
		}
	default:
		responseMsg = "Неизвестный выбор."
		msg := tgbotapi.NewMessage(chatID, responseMsg)

		bot.Send(msg)
	}

	// Обновляем состояние пользователя
	models.UpdateUserState(int64(callback.From.ID), "waiting_selection", userState.Data)

	// Удаление предыдущего сообщения
	bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
		ChatID:    chatID,
		MessageID: msgID,
	})

	// Подтверждение нажатия кнопки
	bot.AnswerCallbackQuery(tgbotapi.NewCallback(callback.ID, "Вы выбрали: "+callback.Data))
}

func SendPhoto(bot *tgbotapi.BotAPI, chatID int64, photoPath string, caption string) error {
	// Открываем файл
	// file, err := os.Open(photoPath)
	// if err != nil {
	// 	return err // Возвращаем ошибку, если файл не найден
	// }
	// defer file.Close() // Закрываем файл после использования

	// Создаем объект PhotoConfig
	photo := tgbotapi.NewPhotoUpload(chatID, photoPath) // Используем NewPhotoUpload
	photo.Caption = caption                             // Добавляем подпись, если необходимо

	// Отправляем фото
	_, err := bot.Send(photo)
	return err
}
