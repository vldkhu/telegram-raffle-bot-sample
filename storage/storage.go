package storage

import "telegram-raffle-bot/models"

var Channels = make(map[int64]*models.Channel) // хранение каналов и их рохыгрышей

func GetChannel(chatID int64) *models.Channel {
	return Channels[chatID]
}

func CreateChannel(chatID int64, ownerID int64) {
	Channels[chatID] = &models.Channel{OwnerID: ownerID, Raffle: &models.Riffle{IsActive: true}}
}
