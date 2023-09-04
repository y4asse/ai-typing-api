package model

import "time"

type User struct {
	UserId    string    `json:"user_id" gorm:"primaryKey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Image     string    `json:"image"`
}

type UserResponse struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	TotalScore int    `json:"total_score"`
}
