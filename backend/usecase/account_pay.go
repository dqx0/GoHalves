package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

func (pu *payUsecase) AddAccountToPay(payId int, accountId int) (model.AccountPay, error) {
	accountPay := model.AccountPay{
		PayID:     uint(payId),
		AccountID: uint(accountId),
	}
	if err := pu.br.GetAccountPayRepository().CreateAccountPay(&accountPay); err != nil {
		return model.AccountPay{}, err
	}
	return accountPay, nil
}
func (pu *payUsecase) DeleteAccountFromPay(payId int, accountId int) (model.AccountPay, error) {
	accountPay := model.AccountPay{}
	atomicBlock := func(br repository.IBaseRepository) error {
		apr := br.GetAccountPayRepository()
		if err := apr.GetAccountPayByPayIdAndAccountId(payId, accountId, &accountPay); err != nil {
			return err
		}
		if err := apr.DeleteAccountPay(int(accountPay.ID), &accountPay); err != nil {
			return err
		}
		return nil
	}
	err := pu.br.Atomic(atomicBlock)
	return accountPay, err
}
