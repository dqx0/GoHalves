package validator

import (
	"github.com/dqx0/GoHalves/go/model"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IEventValidator interface {
	EventValidate(event *model.Event) error
}

type eventValidator struct {
}

func NewEventValidator() IEventValidator {
	return &eventValidator{}
}

func (evv *eventValidator) EventValidate(event *model.Event) error {
	return validation.ValidateStruct(event,
		validation.Field(&event.Title, validation.Required.Error("title is required")),
		validation.Field(&event.Description, validation.Required.Error("description is required")),
	)
}
