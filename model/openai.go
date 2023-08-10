package model

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type AiTextResponse struct {
	Text     []string `json:"text"`
	Hiragana []string `json:"hiragana"`
}

type OpenaiResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Choices []struct {
		Index        int     `json:"index"`
		Message      Message `json:"message"`
		FinishReason string  `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type AiTextRequest struct {
	Thema string `json:"thema"`
}

type AnalyseRequest struct {
	Score         int    `json:"score"`
	Time          int    `json:"time"`
	TypeKeyCount  int    `json:"type_key_count"`
	MissTypeCount int    `json:"miss_type_count"`
	KPM           int    `json:"kpm"`
	Accuracy      int    `json:"accuracy"`
	MissTypeKey   string `json:"miss_type_key"`
}
