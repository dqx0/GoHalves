package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

func (eu *eventUsecase) AddAccountToEvent(eventId int, accountId int, authorityId int) (model.AccountEvent, error) {
	aer := eu.br.GetAccountEventRepository()
	accountEvent := model.AccountEvent{
		AccountID:   uint(accountId),
		EventID:     uint(eventId),
		AuthorityID: uint(authorityId),
	}
	if err := aer.CreateAccountEvent(&accountEvent); err != nil {
		return model.AccountEvent{}, err
	}
	return accountEvent, nil
}

func (eu *eventUsecase) UpdateAuthority(accountEventId int, authorityId int) (model.AccountEvent, error) {
	aer := eu.br.GetAccountEventRepository()
	accountEvent := model.AccountEvent{
		AuthorityID: uint(authorityId),
	}
	if err := aer.UpdateAccountEvent(accountEventId, &accountEvent); err != nil {
		return model.AccountEvent{}, err
	}
	return accountEvent, nil
}
func (eu *eventUsecase) DeleteAccountFromEvent(eventId int, accountId int) (model.AccountEvent, error) {
	accountEvent := model.AccountEvent{}
	atomicBlock := func(br repository.IBaseRepository) error {
		aer := eu.br.GetAccountEventRepository()
		if err := aer.GetAccountEventByEventIdAndAccountId(eventId, accountId, &accountEvent); err != nil {
			return err
		}
		if err := aer.DeleteAccountEvent(int(accountEvent.ID), &accountEvent); err != nil {
			return err
		}
		return nil
	}
	err := eu.br.Atomic(atomicBlock)
	return accountEvent, err
}
