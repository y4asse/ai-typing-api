package model

import "time"

type Batch struct {
	ID         string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name       string    `json:"name"`
	ModeId     int       `json:"mode_id"`
	UserId     string    `json:"user_id"`
	Created_at time.Time `json:"created_at"`
}
