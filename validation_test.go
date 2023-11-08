package validation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidate(t *testing.T) {
	var validate *validator.Validate = validator.New()
	if validate == nil {
		t.Error("validate is nil")
	}
}

func TestValidateVariabel(t *testing.T) {
	validate := validator.New()
	user := "jaja"

	err := validate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidateTwoVariabel(t *testing.T) {
	validate := validator.New()

	password := "secret"
	confirmPassword := "secret"

	err := validate.VarWithValue(password, confirmPassword, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTag(t *testing.T) {
	validate := validator.New()
	user := "12222222222"

	err := validate.Var(user, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestStruct(t *testing.T) {
	type LoginRequest struct {
		Username string `validate:"required,email"`
		Password string `validate:"required,min=5"`
		Age      int    `validate:"gte=0,lte=130"`
	}

	validate := validator.New()
	loginRequest := LoginRequest{
		Username: "jaja@example.com",
		Password: "secret",
		Age:      10,
	}

	err := validate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}
