package model

type Like struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	GameId string `json:"game_id"`
}
