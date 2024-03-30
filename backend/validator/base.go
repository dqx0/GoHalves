package validator

type IBaseValidator interface {
	GetAccountValidator() IAccountValidator
	GetEventValidator() IEventValidator
	GetPayValidator() IPayValidator
}
type baseValidator struct{}

func NewBaseValidator() IBaseValidator {
	return &baseValidator{}
}
func (bv *baseValidator) GetAccountValidator() IAccountValidator {
	return NewAccountValidator()
}
func (bv *baseValidator) GetEventValidator() IEventValidator {
	return NewEventValidator()
}
func (bv *baseValidator) GetPayValidator() IPayValidator {
	return NewPayValidator()
}
