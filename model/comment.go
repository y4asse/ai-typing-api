package model

import "time"

type Comment struct {
	ID           string    `json:"id" gorm:"primaryKey"`
	Name         string    `json:"name"`
	Content      string    `json:"content"`
	CreatedAt    time.Time `json:"created_at"`
	UserId       string    `json:"user_id"`
	PostedTextId string    `json:"posted_text_id"`
}
