package test

import (
	"fmt"
	"testing"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

func TestSimpole(t *testing.T) {
	validate := validator.New()
	var user string = "faridlan"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	validate := validator.New()
	loginReq := LoginRequest{
		Email:    "faridlan@mail.com",
		Password: "rahasia123",
	}

	err := validate.Struct(loginReq)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(loginReq)
	}
}

func TestValidationErrors(t *testing.T) {
	validate := validator.New()
	loginReq := LoginRequest{
		Email:    "faridlan",
		Password: "nul",
	}

	english := english.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(loginReq)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println(fieldError.Translate(trans))
		}
	}
}
