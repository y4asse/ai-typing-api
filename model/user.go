package model

import "time"

type User struct {
	ID         string    `json:"id" gorm:"primaryKey"`
	Name       string    `json:"name"`
	UserId     string    `json:"user_id"`
	TotalScore int       `json:"total_score"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	TotalScore int    `json:"total_score"`
}
