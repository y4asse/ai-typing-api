package model

import "time"

type Like struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	GameId    string    `json:"game_id"`
	CreatedAt time.Time `json:"created_at"`
}

type GameIdCount struct {
	GameId string `json:"game_id"`
	Count  int    `json:"count"`
}
