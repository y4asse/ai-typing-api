package model

type Like struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	PostedTextId string `json:"posted_text_id"`
	UserId       string `json:"user_id"`
}