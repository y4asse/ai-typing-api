package model

type User struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	UserId     string `json:"user_id"`
	TotalScore int    `json:"total_score"`
}

type UserResponse struct {
	ID         string `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	TotalScore int    `json:"total_score"`
}
