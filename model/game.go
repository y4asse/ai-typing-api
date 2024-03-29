package model

import "time"

type Game struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	UserId         string    `json:"user_id"`
	Score          int       `json:"score"`
	InputedThema   string    `json:"inputed_thema"`
	CreatedAt      time.Time `json:"created_at"`
	ModeId         int       `json:"mode_id"`
	TotalKeyCount  int       `json:"total_key_count"`
	TotalMissType  int       `json:"total_miss_type"`
	TotalTime      int       `json:"total_time"`
	DisableRanking bool      `json:"disable_ranking" gorm:"default:false"`
	AiModel        string    `json:"ai_model" gorm:"default:'gpt-3.5-turbo'"`
	Detail         string    `json:"detail" gorm:"default:'についての文章'"`
	MissTypeKeySet string    `json:"miss_type_key_set"`
}

//MissTypeSet=kj,ai,ed => 欲しかったkeyはk,実際に入力されたkeyはj...という意味

type GameResponse struct {
	ID             string    `json:"id" gorm:"primaryKey"`
	UserId         string    `json:"user_id"`
	Score          int       `json:"score"`
	InputedThema   string    `json:"inputed_thema"`
	CreatedAt      time.Time `json:"created_at"`
	ModeId         int       `json:"mode_id"`
	TotalKeyCount  int       `json:"total_key_count"`
	TotalMissType  int       `json:"total_miss_type"`
	TotalTime      int       `json:"total_time"`
	DisableRanking bool      `json:"disable_ranking"`
	AiModel        string    `json:"ai_model"`
	Detail         string    `json:"detail"`
}

type UpdateGameResponse struct {
	Count   int     `json:"count"`
	Rank    int     `json:"rank"`
	Batches []Batch `json:"batches"`
}

type GameDetail struct {
	Game        Game          `json:"game"`
	CreatedText []CreatedText `json:"created_text"`
}
