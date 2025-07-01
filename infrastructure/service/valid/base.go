package valid

import (
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	vi_translations "github.com/go-playground/validator/v10/translations/vi"
)

type Validator struct {
	v     *validator.Validate
	trans ut.Translator
}

func NewValidator(v *validator.Validate) *Validator {
	viLocale := vi.New()
	uni := ut.New(viLocale, viLocale)
	trans, _ := uni.GetTranslator("vi")
	vi_translations.RegisterDefaultTranslations(v, trans)
	return &Validator{v: v, trans: trans}
}

type ValidationError struct {
	Message string
	Data    map[string]any
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (v *Validator) ValidateStruct(s any) *ValidationError {
	err := v.v.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, ve := range validationErrors {
			field := ve.Field()
			errorMessages[field] = ve.Translate(v.trans)
		}
		data := make(map[string]any, len(errorMessages))
		for k, v := range errorMessages {
			data[k] = v
		}
		return &ValidationError{
			Message: "Dữ liệu không hợp lệ",
			Data:    data,
		}
	}
	return nil
}
