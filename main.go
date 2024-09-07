package main

import (
	"log"
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение токена из переменных окружения
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN must be set in .env file")
	}

	// Инициализация бота
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	// bot.Debug = true
	// Создание нового обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// Получение канала обновлений
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	// Обработка обновлений
	for update := range updates {
		if update.Message == nil { // Игнорировать не сообщения
			continue
		}

		// Логирование входящих сообщений
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		// Ответ на команду /start
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет! Я ваш бот.")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Команда не распознана.")
				bot.Send(msg)
			}
		} else {
			// Ответ на текстовые сообщения
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы написали: "+update.Message.Text)
			bot.Send(msg)
		}
	}
}
