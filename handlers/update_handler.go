package handlers

import (
	"log"
	"strconv"
	"telebot/models"
	"telebot/utils"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	lastMessageID int
	lastListID    int
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

		// Обработка команд и нажатий на кнопки
		if update.Message.IsCommand() {
			HandleCommand(bot, update.Message, userState)
		} else {
			HandleButtonPress(bot, update.Message, userState) // Обработка нажатий на кнопки
		}
	} else if update.CallbackQuery != nil { // Обработка нажатий на инлайн кнопки
		handleCallback(bot, update.CallbackQuery)
	}
}

func HandleCommand(bot *tgbotapi.BotAPI, message *tgbotapi.Message, userState *models.UserState) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Выберите год выпуска серии Hot Wheels \n ⬇️⬇️⬇️.")
			replyKeyboard := utils.CreateKeyboard() // Используем постоянную клавиатуру
			msg.ReplyMarkup = replyKeyboard
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

func HandleButtonPress(bot *tgbotapi.BotAPI, message *tgbotapi.Message, userState *models.UserState) {
	chatID := message.Chat.ID

	if messageIDs, ok := userState.Data["sent_message_ids"].([]int); ok && len(messageIDs) > 0 {
		for _, msgID := range messageIDs {
			// Попробуем удалить сообщение
			if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
				ChatID:    chatID,
				MessageID: msgID,
			}); err != nil {
				log.Printf("Ошибка при удалении сообщения с ID %d: %s", msgID, err)
			}
		}
		// Обнуляем слайс после удаления всех сообщений
		userState.Data["sent_message_ids"] = []int{}
	}

	if lastListID != 0 {
		if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: lastListID,
		}); err != nil {
			log.Printf("Ошибка при удалении сообщения: %s", err)
		}
	}
	if lastMessageID != 0 {
		if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: lastMessageID,
		}); err != nil {
			log.Printf("Ошибка при удалении сообщения: %s", err)
		}
	}

	var responseMsg string

	// Обработка нажатий на кнопки
	switch message.Text {
	case "2023-TH":
		userState.Data["selected_year"] = "2023"
		userState.Data["selected_series"] = "th"
		responseMsg = getSelectedMessage("2023", "th")
	case "2023-Supers":
		userState.Data["selected_year"] = "2023"
		userState.Data["selected_series"] = "supers"
		responseMsg = getSelectedMessage("2023", "supers")
	case "2024-TH":
		userState.Data["selected_year"] = "2024"
		userState.Data["selected_series"] = "th"
		responseMsg = getSelectedMessage("2024", "th")
	case "2024-Supers":
		userState.Data["selected_year"] = "2024"
		userState.Data["selected_series"] = "supers"
		responseMsg = getSelectedMessage("2024", "supers")
	default:
		responseMsg = "Непонятный выбор."
	}

	// Создаем сообщение для отправки
	msg := tgbotapi.NewMessage(chatID, responseMsg)
	msg.ReplyMarkup = utils.CreateKeyboardPhotos()
	// Отправляем сообщение
	sentMsg, err := bot.Send(msg)
	if err == nil {
		lastListID = sentMsg.MessageID // Сохраняем ID отправленного сообщения
	}
}

func handleCallback(bot *tgbotapi.BotAPI, callback *tgbotapi.CallbackQuery) {
	log.Printf("Кнопка нажата: %s", callback.Data)
	chatID := callback.Message.Chat.ID
	userState := models.GetUserState(int64(callback.From.ID))
	responseMsg := "Неизвестный выбор."

	year, yearOk := userState.Data["selected_year"].(string)
	series, seriesOk := userState.Data["selected_series"].(string)

	if !yearOk || year == "" || !seriesOk || series == "" {
		bot.Send(tgbotapi.NewMessage(chatID, responseMsg))
		return
	}

	switch callback.Data {
	case "all":
		if year, ok := userState.Data["selected_year"].(string); ok && year != "" {
			if series, ok := userState.Data["selected_series"].(string); ok && series != "" {
				messageIDs := sendAllPhotos(bot, chatID, year, series)

				// Сохраняем messageIDs в состояние пользователя, если нужно
				userState.Data["sent_message_ids"] = messageIDs
			}
		}
	default:
		if num, err := strconv.Atoi(callback.Data); err == nil {
			sendSinglePhoto(bot, chatID, year, series, num)
		} else {
			bot.Send(tgbotapi.NewMessage(chatID, responseMsg))
		}
	}
}

func sendSinglePhoto(bot *tgbotapi.BotAPI, chatID int64, year, series string, number int) {
	file := getPhoto(year, series, number)

	if file != nil {
		for key, value := range file.(map[string]string) {
			media := tgbotapi.NewPhotoUpload(chatID, value)
			media.Caption = key // Используем ключ как подпись

			if lastMessageID != 0 {
				if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
					ChatID:    chatID,
					MessageID: lastMessageID,
				}); err != nil {
					log.Printf("Ошибка при удалении сообщения: %s", err)
				}
			}

			sentMsg, err := bot.Send(media)
			if err != nil {
				log.Printf("Ошибка при отправке фото %s: %s", value, err)
			} else {
				lastMessageID = sentMsg.MessageID // Сохраняем ID нового сообщения
			}
		}
	}
}

func sendAllPhotos(bot *tgbotapi.BotAPI, chatID int64, year, series string) []int {
	files := getPhoto(year, series, 0) // Получаем все файлы
	var messageIDs []int               // Слайс для хранения ID отправленных сообщений

	if lastMessageID != 0 {
		if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
			ChatID:    chatID,
			MessageID: lastMessageID,
		}); err != nil {
			log.Printf("Ошибка при удалении сообщения: %s", err)
		}
	}
	// if lastListID != 0 {
	// 	if _, err := bot.DeleteMessage(tgbotapi.DeleteMessageConfig{
	// 		ChatID:    chatID,
	// 		MessageID: lastListID,
	// 	}); err != nil {
	// 		log.Printf("Ошибка при удалении сообщения: %s", err)
	// 	}
	// }
	if files != nil {
		for key, value := range files.(map[string]string) {
			media := tgbotapi.NewPhotoUpload(chatID, value)
			media.Caption = key // Используем ключ как подпись

			sentMsg, err := bot.Send(media)
			if err != nil {
				log.Printf("Ошибка при отправке фото %s: %s", value, err)
				continue // Пропускаем, если произошла ошибка
			}
			messageIDs = append(messageIDs, sentMsg.MessageID) // Сохраняем ID отправленного сообщения
		}
	}
	return messageIDs // Возвращаем слайс с ID
}

func SendPhoto(bot *tgbotapi.BotAPI, chatID int64, photoPath string, caption string) error {
	photo := tgbotapi.NewPhotoUpload(chatID, photoPath) // Используем NewPhotoUpload
	photo.Caption = caption                             // Добавляем подпись, если необходимо

	_, err := bot.Send(photo)
	return err
}
