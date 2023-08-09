package model

type GameWithCount struct {
	Game  Game `json:"game"`
	Count int  `json:"count"`
}
