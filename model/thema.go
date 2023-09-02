package model

type ThemaRequest struct {
	Thema   string `json:"thema"`
	AiModel string `json:"aiModel"`
	Detail  string `json:"detail"`
}
