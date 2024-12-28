package bot

import (
	"fmt"
	"log"
	"telegram-raffle-bot/models"
	"telegram-raffle-bot/raffle"
	"telegram-raffle-bot/storage"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	updateConfig := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Добро пожаловать! Используйте /create_raffle для создания розыгрыша.")
			bot.Send(msg)

		case "/create_raffle":
			storage.CreateChannel(update.Message.Chat.ID, update.Message.From.ID)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Розыгрыш создан! Используйте /register для регистрации участников.")
			bot.Send(msg)

		case "/register":
			participant := models.Participant{ID: update.Message.From.ID, Name: update.Message.From.UserName}
			raffle.RegisterParticipant(update.Message.Chat.ID, participant)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Вы зарегистрированы!")
			bot.Send(msg)

		case "/set_time":
			// Пример установки времени розыгрыша (здесь можно добавить парсинг времени из сообщения)
			drawTime := time.Now().Add(1 * time.Minute).Unix() // Пример: через 1 минуту
			raffle.SetDrawTime(update.Message.Chat.ID, drawTime)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Время розыгрыша установлено на: %s", time.Unix(drawTime, 0).Format(time.RFC1123)))
			bot.Send(msg)

		case "/draw":
			winner, err := raffle.DrawWinner(update.Message.Chat.ID)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
				bot.Send(msg)
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Победитель: %s", winner))
				bot.Send(msg)

			}

		}
	}
}
