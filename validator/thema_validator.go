package validator

import (
	"ai-typing/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IAiTextValidator interface {
	ThemaValidator(model.ThemaRequest) error
}

type aiTextValidator struct{}

func NewAitextValidator() IAiTextValidator {
	return &aiTextValidator{}
}

func (aiTextValidator *aiTextValidator) ThemaValidator(themaRequest model.ThemaRequest) error {
	return validation.ValidateStruct(&themaRequest,
		validation.Field(
			&themaRequest.Thema,
			validation.Required.Error("テーマは必須です"),
			validation.RuneLength(1, 10).Error("テーマは10文字以内で入力してください"),
		),
	)
}
