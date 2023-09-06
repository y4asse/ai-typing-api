package model

import "time"

type Batch struct {
	ID        string    `json:"id" gorm:"primaryKey;default:uuid_generate_v4()"`
	Name      string    `json:"name" gorm:"uniqueIndex:user_id_name_uniq"`
	ModeId    int       `json:"mode_id" gorm:"uniqueIndex:user_id_name_uniq"`
	UserId    string    `json:"user_id" gorm:"uniqueIndex:user_id_name_uniq"`
	CreatedAt time.Time `json:"created_at"`
}
