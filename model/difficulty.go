package model

type Difficulty struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Difficulty string `json:"difficulty"`
}
