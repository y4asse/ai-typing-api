package model

import "time"

type Game struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	UserId        string    `json:"user_id"`
	Score         int       `json:"score"`
	InputedThema  string    `json:"inputed_thema"`
	CreatedAt     time.Time `json:"created_at"`
	ModeId        int       `json:"mode_id"`
	TotalKeyCount int       `json:"total_key_count"`
	TotalMissType int       `json:"total_miss_type"`
	TotalTime     int       `json:"total_time"`
}

type GameResponse struct {
	ID            string    `json:"id" gorm:"primaryKey"`
	UserId        string    `json:"user_id"`
	Score         int       `json:"score"`
	InputedThema  string    `json:"inputed_thema"`
	CreatedAt     time.Time `json:"created_at"`
	ModeId        int       `json:"mode_id"`
	TotalKeyCount int       `json:"total_key_count"`
	TotalMissType int       `json:"total_miss_type"`
	TotalTime     int       `json:"total_time"`
}
