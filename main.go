package main

import (
	"log"
	"telebot/config"
	"telebot/database"
	"telebot/handlers"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

	// Инициализация базы данных
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Error initializing database: ", err)
	}
	defer db.Close()

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

		// Обработка сообщений и команд
		handlers.HandleMessage(db, bot, update.Message)
	}
}
