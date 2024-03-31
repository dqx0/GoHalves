package controller

import (
	"github.com/dqx0/GoHalves/go/usecase"
)

type IBaseController interface {
	GetAccountController() IAccountController
}
type baseController struct {
	bu usecase.IBaseUsecase
	su usecase.ISessionUsecase
}

func NewBaseController(bu usecase.IBaseUsecase, su usecase.ISessionUsecase) IBaseController {
	return &baseController{bu, su}
}
func (bc *baseController) GetAccountController() IAccountController {
	return NewAccountController(bc.bu, bc.su)
}
