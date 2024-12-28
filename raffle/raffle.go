package raffle

import (
	"fmt"
	"math/rand"
	"telegram-raffle-bot/models"
	"telegram-raffle-bot/storage"
	"time"
)

func RegisterParticipant(chatID int64, participant models.Participant) {
	channel := storage.GetChannel(chatID)
	if channel != nil && channel.Raffle.IsActive {
		channel.Raffle.Participants = append(channel.Raffle.Participants, participant)
	}

}

func SetDrawTime(chatID int64, drawTime int64) {
	channel := storage.GetChannel(chatID)
	if channel != nil {
		channel.Raffle.DrawTime = drawTime
	}
}

func DrawWinner(chatID int64) (string, error) {
	channel := storage.GetChannel(chatID)
	if channel != nil && channel.Raffle.IsActive {
		if len(channel.Raffle.Participants) == 0 {
			return "", fmt.Errorf("нет зарегистрированных пользователей")
		}
		rand.NewSource(time.Now().UnixNano())
		winner := channel.Raffle.Participants[rand.Intn(len(channel.Raffle.Participants))]
		channel.Raffle.IsActive = false
		return winner.Name, nil
	}
	return "", fmt.Errorf("нет активного розыгрыша")
}
