package main

import (
	"log"
	"telebot/config"
	"telebot/utils"
	//	"telebot/database"
	"telebot/handlers"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// // Инициализация базы данных
	// db, err := database.InitDB(cfg.DatabaseURL)
	// if err != nil {
	// 	log.Fatal("Error initializing database: ", err)
	// }
	// defer db.Close()

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal("Error initializing bot: ", err)
	}
	bot.Debug = cfg.Debug

	// Обработка обновлений
	updates := handlers.SetupUpdates(bot)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		// Определите кнопки
		btn1 := tgbotapi.NewKeyboardButton("Команда 1")
		btn2 := tgbotapi.NewKeyboardButton("Команда 2")
		btn3 := tgbotapi.NewKeyboardButton("Команда 3")
		row := []tgbotapi.KeyboardButton{btn1, btn2, btn3}
		// Создайте клавиатуру
		keyboard1 := tgbotapi.NewReplyKeyboard(row)

		// Создаем клавиатуру
		keyboard := utils.CreateKeyboardYear()

		// Отправляем сообщение с клавиатурой
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери год выпука серии:")
		msg.ReplyMarkup = keyboard
		msg.ReplyMarkup = keyboard1

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
		// Обработка сообщений и команд
		handlers.HandleMessage(bot, update.Message)
		// //Обработка сообщений и команд c database
		// handlers.HandleMessageDB(db, bot, update.Message)

	}
}
