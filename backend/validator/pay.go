package validator

import (
	"github.com/dqx0/GoHalves/go/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IPayValidator interface {
	PayValidate(pay *model.Pay) error
	CreatePayValidate(pay *model.Pay) error
}

type payValidator struct {
}

func NewPayValidator() IPayValidator {
	return &payValidator{}
}

func (pv *payValidator) PayValidate(pay *model.Pay) error {
	return validation.ValidateStruct(pay,
		validation.Field(&pay.ID, validation.Required.Error("id is required")),
		validation.Field(&pay.EventID, validation.Required.Error("event_id is required")),
		validation.Field(&pay.Amount, validation.Required.Error("amount is required"), validation.Min(1).Error("amount must be greater than 0")),
	)
}
func (pv *payValidator) CreatePayValidate(pay *model.Pay) error {
	return validation.ValidateStruct(pay,
		validation.Field(&pay.EventID, validation.Required.Error("event_id is required")),
		validation.Field(&pay.Amount, validation.Required.Error("amount is required"), validation.Min(1).Error("amount must be greater than 0")),
	)
}
