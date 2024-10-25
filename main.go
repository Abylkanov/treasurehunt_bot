package main

import (
	"log"
	"telebot/config"
	//"telebot/utils"
	"telebot/handlers"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var status = "root"

func main() {
	// Загрузка конфигурации
	cfg := config.LoadConfig()

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

	}
}
