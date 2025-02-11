package helper

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"reflect"
)

func SetCustomMessageValidator(validate *validator.Validate, trans *ut.Translator) {
	_ = validate.RegisterTranslation("startswith", *trans, func(ut ut.Translator) error {
		return ut.Add("startswith", "{0} harus berawalan {1} !", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		fieldName := fe.Field() // Nama field
		param := fe.Param()     // Nilai parameter (misalnya '8')
		t, _ := ut.T("startswith", fieldName, param)
		return t
	})

}

func SetLanguageID() (*validator.Validate, *ut.Translator) {
	Indonesia := id.New()
	uni := ut.New(Indonesia, Indonesia)
	trans, _ := uni.GetTranslator("id")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		alias := fld.Tag.Get("json")
		if alias == "" {
			return fld.Name
		}
		return alias
	})

	SetCustomMessageValidator(validate, &trans)

	err := idTranslations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic("register translation error")
	}

	return validate, &trans
}
