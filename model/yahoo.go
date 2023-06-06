package model

type YahooRes struct {
	ID      string `json:"id"`
	Jsonrpc string `json:"jsonrpc"`
	Result  Result `json:"result"`
}

type Result struct {
	Word []Word `json:"word"`
}

type Word struct {
	Surface  string    `json:"surface"`
	Furigana string    `json:"furigana"`
	Roman    string    `json:"roman"`
	Subword  []Subword `json:"subword"`
}

type Subword struct {
	Surface  string `json:"surface"`
	Furigana string `json:"furigana"`
	Roman    string `json:"roman"`
}
