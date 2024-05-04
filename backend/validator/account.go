package validator

import (
	"errors"
	"regexp"

	"github.com/dqx0/GoHalves/go/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IAccountValidator interface {
	AccountValidate(account *model.Account) error
	UpdateAccountValidate(account *model.Account) error
}
type accountValidator struct{}

func NewAccountValidator() IAccountValidator {
	return &accountValidator{}
}

func (acv *accountValidator) AccountValidate(account *model.Account) error {
	return validation.ValidateStruct(account,
		validation.Field(&account.UserID, validation.Required.Error("user_id is required")),
		validation.Field(&account.Name, validation.Required.Error("name is required")),
		validation.Field(&account.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(4, 20).Error("password must be at 4 ~ 20 characters"),
		),
	)
}
func (acv *accountValidator) UpdateAccountValidate(account *model.Account) error {
	return validation.ValidateStruct(account,
		validation.Field(&account.UserID, validation.Required.Error("user_id is required")),
		validation.Field(&account.Name, validation.Required.Error("name is required")),
		validation.Field(&account.Email, validation.By(func(value interface{}) error {
			return acv.isEmail(value.(string))
		})),
	)
}
func (acv *accountValidator) isEmail(value string) error {
	emailRegex := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	match := regexp.MustCompile(emailRegex).MatchString
	if !match(value) {
		return errors.New("email must be a valid format")
	}
	return nil
}
