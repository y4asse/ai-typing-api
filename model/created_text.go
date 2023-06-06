package model

type CreatedText struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	GamaId  string `json:"game_id"`
	IsPost  bool   `json:"is_post"`
}
type CreatedTextResponse struct {
	ID      string `json:"id" gorm:"primaryKey"`
	Content string `json:"content"`
	GamaId  string `json:"game_id"`
	IsPost  bool   `json:"is_post"`
}


