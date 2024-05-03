package handler

import (
	"github.com/dqx0/GoHalves/go/usecase"
)

type IBaseHandler interface {
	GetAccountHandler() IAccountHandler
	GetEventHandler() IEventHandler
	GetPayHandler() IPayHandler
	GetFriendHandler() IFriendHandler
	GetSessionHandler() ISessionHandler
	GetTestHandler() ITestHandler
}
type baseHandler struct {
	bu usecase.IBaseUsecase
}

func NewBaseHandler(bu usecase.IBaseUsecase) IBaseHandler {
	return &baseHandler{bu}
}
func (bc *baseHandler) GetAccountHandler() IAccountHandler {
	return NewAccountHandler(bc.bu)
}
func (bc *baseHandler) GetEventHandler() IEventHandler {
	return NewEventHandler(bc.bu)
}
func (bc *baseHandler) GetPayHandler() IPayHandler {
	return NewPayHandler(bc.bu)
}
func (bc *baseHandler) GetFriendHandler() IFriendHandler {
	return NewFriendHandler(bc.bu)
}
func (bc *baseHandler) GetSessionHandler() ISessionHandler {
	return NewSessionHandler(bc.bu)
}
func (bc *baseHandler) GetTestHandler() ITestHandler {
	return NewTestHandler()
}
