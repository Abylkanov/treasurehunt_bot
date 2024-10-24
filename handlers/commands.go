package handlers

import (
	//"database/sql"
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func HandleMessage(bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
	// Логирование входящих сообщений
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	// Ответ на команду /start
	if message.IsCommand() {
		switch message.Command() {
		case "start":
			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Выберите из меню год выпуска серии Hot Wheels \n ⬇️⬇️⬇️.")
			bot.Send(msg)
		default:
			msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не распознана.")
			bot.Send(msg)
		}
	} else {
		// Ответ на текстовые сообщения
		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы написали: "+message.Text)
		bot.Send(msg)
	}
}

// func HandleMessageDB(db *sql.DB, bot *tgbotapi.BotAPI, message *tgbotapi.Message) {
// 	// Логирование входящих сообщений
// 	log.Printf("[%s] %s", message.From.UserName, message.Text)

// 	// Сохранение пользователя в базе данных
// 	_, err := db.Exec(`INSERT INTO users (id, username, first_name, last_name) VALUES ($1, $2, $3, $4)
//         ON CONFLICT (id) DO UPDATE SET username = EXCLUDED.username, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name`,
// 		message.From.ID,
// 		message.From.UserName,
// 		message.From.FirstName,
// 		message.From.LastName)
// 	if err != nil {
// 		log.Printf("Error saving user to database: %v", err)
// 	}

// 	// Ответ на команду /start
// 	if message.IsCommand() {
// 		switch message.Command() {
// 		case "start":
// 			msg := tgbotapi.NewMessage(message.Chat.ID, "Привет! Я ваш бот.")
// 			bot.Send(msg)
// 		default:
// 			msg := tgbotapi.NewMessage(message.Chat.ID, "Команда не распознана.")
// 			bot.Send(msg)
// 		}
// 	} else {
// 		// Ответ на текстовые сообщения
// 		msg := tgbotapi.NewMessage(message.Chat.ID, "Вы написали: "+message.Text)
// 		bot.Send(msg)
// 	}
// }
