package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/dqx0/GoHalves/go/validator"
)

type IPayUsecase interface {
	GetPaysByEventId(eventId int) ([]model.Pay, error)
	GetPayById(id int) (model.Pay, error)
	GetPaysByAccountIdAndEventId(accountId int, eventId int) ([]model.Pay, error)
	CreatePay(pay model.Pay, createdAccountId int, accountIdsToPay []int) (model.Pay, error)
	AddAccountToPay(payId int, accountId int) (model.AccountPay, error)
	UpdatePay(id int, pay model.Pay) (model.Pay, error)
	DeletePay(id int) (model.Pay, error)
	DeleteAccountFromPay(payId int, accountId int) (model.AccountPay, error)
}
type payUsecase struct {
	br repository.IBaseRepository
	bv validator.IBaseValidator
}

func NewPayUsecase(br repository.IBaseRepository, bv validator.IBaseValidator) IPayUsecase {
	return &payUsecase{br, bv}
}
func (pu *payUsecase) GetPaysByEventId(eventId int) ([]model.Pay, error) {
	pays := []model.Pay{}
	pr := pu.br.GetPayRepository()
	if err := pr.GetPaysByEventId(eventId, &pays); err != nil {
		return nil, err
	}
	return pays, nil
}
func (pu *payUsecase) GetPayById(id int) (model.Pay, error) {
	pay := model.Pay{}
	pr := pu.br.GetPayRepository()
	if err := pr.GetPayById(id, &pay); err != nil {
		return model.Pay{}, err
	}
	return pay, nil
}
func (pu *payUsecase) GetPaysByAccountIdAndEventId(accountId int, eventId int) ([]model.Pay, error) {
	pays := []model.Pay{}
	pr := pu.br.GetPayRepository()
	if err := pr.GetPaysByAccountIdAndEventId(accountId, eventId, &pays); err != nil {
		return nil, err
	}
	return pays, nil
}
func (pu *payUsecase) CreatePay(pay model.Pay, createdAccountId int, accountIdsToPay []int) (model.Pay, error) {
	pv := pu.bv.GetPayValidator()

	// バリデーション
	if err := pv.PayValidate(&pay); err != nil {
		return model.Pay{}, err
	}

	atomicBlock := func(br repository.IBaseRepository) error {
		pr := br.GetPayRepository()

		// 支払い作成
		if err := pr.CreatePay(&pay); err != nil {
			return err
		}

		for _, accountId := range accountIdsToPay {
			accountPay := model.AccountPay{
				AccountID: uint(accountId),
				PayID:     pay.ID,
			}

			// アカウント支払い作成
			apr := pu.br.GetAccountPayRepository()
			if err := apr.CreateAccountPay(&accountPay); err != nil {
				if err := pr.DeletePay(int(pay.ID), &pay); err != nil {
					return err
				}
				return err
			}
		}
		return nil
	}

	err := pu.br.Atomic(atomicBlock)
	return pay, err
}
func (pu *payUsecase) UpdatePay(id int, pay model.Pay) (model.Pay, error) {
	pr := pu.br.GetPayRepository()
	if err := pr.UpdatePay(id, &pay); err != nil {
		return model.Pay{}, err
	}
	return pay, nil
}
func (pu *payUsecase) DeletePay(id int) (model.Pay, error) {
	pr := pu.br.GetPayRepository()
	pay := model.Pay{}
	err := pr.GetPayById(id, &pay)
	if err != nil {
		return model.Pay{}, err
	}
	if err := pr.DeletePay(id, &pay); err != nil {
		return model.Pay{}, err
	}
	return pay, nil
}
