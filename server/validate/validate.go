package validate

import (
	"github.com/Puttipong1/assessment-tax/config"
	"github.com/Puttipong1/assessment-tax/model/response"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	trans    ut.Translator
	validate *validator.Validate
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return err
	}
	return nil
}
func New() *CustomValidator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New(validator.WithRequiredStructEnabled())
	en_translations.RegisterDefaultTranslations(validate, trans)
	return &CustomValidator{Validator: validate}
}

func ErrorMessage(err error) []response.Error {
	log := config.Logger()
	log.Error().Msg(err.Error())
	errs := err.(validator.ValidationErrors)
	errors := []response.Error{}
	for _, val := range errs.Translate(trans) {
		errors = append(errors, response.Error{
			Message: val,
		})
	}
	return errors
}
