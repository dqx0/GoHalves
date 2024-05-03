package usecase

import (
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/dqx0/GoHalves/go/validator"
)

type IBaseUsecase interface {
	GetAccountUsecase() IAccountUsecase
	GetEventUsecase() IEventUsecase
	GetPayUsecase() IPayUsecase
	GetCalcUsecase() ICalcUsecase
	GetAuthorityUsecase() IAuthorityUsecase
	GetFriendUsecase() IFriendUsecase
	GetSessionUsecase() ISessionUsecase
}
type baseUsecase struct {
	br repository.IBaseRepository
	bv validator.IBaseValidator
}

func NewBaseUsecase(br repository.IBaseRepository, bv validator.IBaseValidator) IBaseUsecase {
	return &baseUsecase{br, bv}
}
func (bu *baseUsecase) GetAccountUsecase() IAccountUsecase {
	return NewAccountUsecase(bu.br, bu.bv)
}
func (bu *baseUsecase) GetEventUsecase() IEventUsecase {
	return NewEventUsecase(bu.br, bu.bv)
}
func (bu *baseUsecase) GetPayUsecase() IPayUsecase {
	return NewPayUsecase(bu.br, bu.bv)
}
func (bu *baseUsecase) GetCalcUsecase() ICalcUsecase {
	return NewCalcUsecase(bu.br)
}
func (bu *baseUsecase) GetAuthorityUsecase() IAuthorityUsecase {
	return NewAuthorityUsecase(bu.br)
}
func (bu *baseUsecase) GetFriendUsecase() IFriendUsecase {
	return NewFriendUsecase(bu.br)
}
func (bu *baseUsecase) GetSessionUsecase() ISessionUsecase {
	return NewSessionUsecase(bu.br)
}
