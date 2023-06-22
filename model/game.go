package model

import "time"

type Game struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	UserId       string    `json:"user_id"`
	Score        int       `json:"score"`
	InputedThema string    `json:"inputed_thema"`
	CreatedAt    time.Time `json:"created_at"`
	ModeId       int       `json:"mode_id"`
}

type GameResponse struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	UserId       string    `json:"user_id"`
	Score        int       `json:"score"`
	InputedThema string    `json:"inputed_thema"`
	CreatedAt    time.Time `json:"created_at"`
	ModeId       int       `json:"mode_id"`
}
