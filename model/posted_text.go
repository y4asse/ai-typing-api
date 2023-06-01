package model

type PostedText struct {
	ID            string `json:"id" gorm:"primaryKey"`
	CreatedTextId string `json:"created_text_id"`
	Comment       string `json:"comment"`
}

type PostedTextResponse struct {
	ID            string `json:"id" gorm:"primaryKey"`
	CreatedTextId string `json:"created_text_id"`
	Comment       string `json:"comment"`
}


