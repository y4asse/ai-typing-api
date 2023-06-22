package model

type TempGame struct {
	UserId       string `json:"user_id"`
	Score        int    `json:"score"`
	InputedThema string `json:"inputed_thema"`
	ModeId       int    `json:"mode_id"`
	Text         string `json:"text"`
	Hiragana     string `json:"hiragana"`
}
