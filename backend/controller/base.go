package controller

import (
	"github.com/dqx0/GoHalves/go/usecase"
)

type IBaseController interface {
	GetAccountController() IAccountController
}
type baseController struct {
	bu usecase.IBaseUsecase
}

func NewBaseController(bu usecase.IBaseUsecase) IBaseController {
	return &baseController{bu}
}
func (bc *baseController) GetAccountController() IAccountController {
	return NewAccountController(bc.bu.GetAccountUsecase())
}
