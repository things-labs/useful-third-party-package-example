package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/zh"
)

// use a single instance , it caches struct info
var validate *validator.Validate
var trans ut.Translator

func main() {
	validate = validator.New()

	// NOTE: ommitting allot of error checking for brevity

	zhx := zh.New()
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ = ut.New(zhx, zhx).GetTranslator("zh")

	en_translations.RegisterDefaultTranslations(validate, trans)

	translateAll(trans)
	translateIndividual(trans)
	translateOverride(trans) // yep you can specify your own in whatever locale you want!
}

func translateAll(trans ut.Translator) {
	type User struct {
		Username string `validate:"required"`
		Tagline  string `validate:"required,lt=10"`
		Tagline2 string `validate:"required,gt=1"`
	}

	user := User{
		Username: "Joeybloggs",
		Tagline:  "This tagline is way too long.",
		Tagline2: "1",
	}

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(TranslatorErr(err))
	}
}

func translateIndividual(trans ut.Translator) {
	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(TranslatorErr(err))
	}
}

func translateOverride(trans ut.Translator) {
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	type User struct {
		Username string `validate:"required"`
	}

	var user User

	err := validate.Struct(user)
	if err != nil {
		fmt.Println(TranslatorErr(err))
	}
}

func TranslatorErr(err error) error {
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		buff := new(bytes.Buffer)
		for i := 0; i < len(errs); i++ {
			// can translate each error one at a time.
			buff.WriteString(errs[i].Translate(trans))
			if i != len(errs)-1 {
				buff.WriteString(", ")
			}
		}
		err = errors.New(strings.TrimSpace(buff.String()))
	}
	return err
}
