package model

type CreatedText struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Text     string `json:"text"`
	Hiragana string `json:"hiragana"`
	GameId   string `json:"game_id" gorm:"foregignKey:GameId"`
	IsPost   bool   `json:"is_post"`
}
type CreatedTextResponse struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Text     string `json:"text"`
	Hiragana string `json:"hiragana"`
	GameId   string `json:"game_id"`
	IsPost   bool   `json:"is_post"`
}
