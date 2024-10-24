package main

import (
	"log"
	"telebot/config"
	//"telebot/utils"
	//	"telebot/database"
	"telebot/handlers"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var status = "root"

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
		switch status {
		case "root":
			status = handlers.HandleCommand(bot, &update)
		case "year":
			handlers.HandleMessage(bot, &update)

		default:

		}
		// Обработка сообщений и команд

		// //Обработка сообщений и команд c database
		// handlers.HandleMessageDB(db, bot, update.Message)

	}
}
