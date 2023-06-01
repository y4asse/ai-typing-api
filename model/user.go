package model

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	UserId       string    `json:"user_id"`
	TotalScore   int       `json:"total_score"`
	Email        string    `json:"email" gorm:"unique"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserResponse struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	TotalScore   int       `json:"total_score"`
	Email        string    `json:"email" gorm:"unique"`
}