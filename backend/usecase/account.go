package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/dqx0/GoHalves/go/validator"
)

type IAccountUsecase interface {
	GetAccounts() ([]model.Account, error)
	GetAccountById(accountId int) (model.Account, error)
	CreateAccount(account model.Account) (model.Account, error)
	UpdateAccount(accountId int, account model.Account) (model.Account, error)
	DeleteAccount(accountId int, account model.Account) (model.Account, error)
}
type accountUsecase struct {
	ar repository.IAccountRepository
	av validator.IAccountValidator
}

func NewAccountUsecase(ar repository.IAccountRepository, av validator.IAccountValidator) IAccountUsecase {
	return &accountUsecase{ar, av}
}
func (au *accountUsecase) GetAccounts() ([]model.Account, error) {
	accounts := []model.Account{}
	if err := au.ar.GetAccounts(&accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}
func (au *accountUsecase) GetAccountById(accountId int) (model.Account, error) {
	account := model.Account{}
	if err := au.ar.GetAccountById(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) CreateAccount(account model.Account) (model.Account, error) {
	if err := au.av.AccountValidate(&account); err != nil {
		return model.Account{}, err
	}
	if err := au.ar.CreateAccount(&account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) UpdateAccount(accountId int, account model.Account) (model.Account, error) {
	if err := au.av.AccountValidate(&account); err != nil {
		return model.Account{}, err
	}
	if err := au.ar.UpdateAccount(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) DeleteAccount(accountId int, account model.Account) (model.Account, error) {
	if err := au.ar.DeleteAccount(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
