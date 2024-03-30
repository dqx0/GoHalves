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
	br repository.IBaseRepository
	bv validator.IBaseValidator
}

func NewAccountUsecase(br repository.IBaseRepository, bv validator.IBaseValidator) IAccountUsecase {
	return &accountUsecase{br, bv}
}
func (au *accountUsecase) GetAccounts() ([]model.Account, error) {
	accounts := []model.Account{}
	ar := au.br.GetAccountRepository()
	if err := ar.GetAccounts(&accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}
func (au *accountUsecase) GetAccountById(accountId int) (model.Account, error) {
	account := model.Account{}
	ar := au.br.GetAccountRepository()
	if err := ar.GetAccountById(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) CreateAccount(account model.Account) (model.Account, error) {
	av := au.bv.GetAccountValidator()
	ar := au.br.GetAccountRepository()
	if err := av.AccountValidate(&account); err != nil {
		return model.Account{}, err
	}
	if err := ar.CreateAccount(&account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) UpdateAccount(accountId int, account model.Account) (model.Account, error) {
	av := au.bv.GetAccountValidator()
	ar := au.br.GetAccountRepository()
	if err := av.AccountValidate(&account); err != nil {
		return model.Account{}, err
	}
	if err := ar.UpdateAccount(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
func (au *accountUsecase) DeleteAccount(accountId int, account model.Account) (model.Account, error) {
	ar := au.br.GetAccountRepository()
	if err := ar.DeleteAccount(accountId, &account); err != nil {
		return model.Account{}, err
	}
	return account, nil
}
