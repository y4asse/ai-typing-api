package model

type Mode struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Mode string `json:"mode"`
}
