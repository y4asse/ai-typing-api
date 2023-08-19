package model

type GameBody struct {
	UserId        string   `json:"user_id"`
	Score         int      `json:"score"`
	InputedThema  string   `json:"inputed_thema"`
	ModeId        int      `json:"mode_id"`
	Text          []string `json:"text"`
	Hiragana      []string `json:"hiragana"`
	TotalKeyCount int      `json:"total_key_count"`
	TotalMissType int      `json:"total_miss_type"`
	TotalTime     int      `json:"total_time"`
}
