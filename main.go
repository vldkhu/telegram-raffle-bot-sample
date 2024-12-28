package main

import (
	"log"
	"os"
	"telegram-raffle-bot/bot"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot.StartBot(token)
}
