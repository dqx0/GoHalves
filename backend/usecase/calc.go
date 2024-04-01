package usecase

import (
	"fmt"
	"math"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type ICalcUsecase interface {
	CalculatePaymentForAccounts(eventId int) ([]model.Calc, error)
}
type calcUsecase struct {
	br repository.IBaseRepository
}

func NewCalcUsecase(br repository.IBaseRepository) ICalcUsecase {
	return &calcUsecase{br}
}
func (cu *calcUsecase) CalculatePaymentForAccounts(eventId int) ([]model.Calc, error) {
	event := model.Event{}
	er := cu.br.GetEventRepository()
	if err := er.GetEventById(eventId, &event); err != nil {
		return nil, err
	}
	calcs := []model.Calc{}
	atomicBlock := func(br repository.IBaseRepository) error {
		pr := br.GetPayRepository()
		aer := br.GetAccountEventRepository()
		apr := br.GetAccountPayRepository()

		// 支払い情報を取得
		if err := pr.GetPaysByEventId(int(event.ID), &event.Pays); err != nil {
			return err
		}

		// イベントに参加しているアカウントを取得
		if err := aer.GetAccountsByEventId(int(event.ID), &event.Accounts); err != nil {
			return err
		}

		// 互いの支払い金額を計算するための構造体を作成
		for _, account := range event.Accounts {
			calc := model.Calc{
				AccountId:      int(account.ID),
				EventId:        int(event.ID),
				AccountAmounts: make(map[int]int),
			}
			for _, otherAccount := range event.Accounts {
				if otherAccount.ID != account.ID {
					calc.AccountAmounts[int(otherAccount.ID)] = 0
				}
			}
			calcs = append(calcs, calc)
		}
		// 支払いごとの対象者をaccount_payから、立替者をpay.PaidUserIDから、
		// 金額をpay.Amountから取得し、金額を対象者の数で割って小数第1位で四捨五入
		// する。AccountIdが対象者のID、AccountAmountsのキーが立替者のID、値が金額
		// となるようにAccountAmountsを更新する。AccountIdが立替者のIDと一致する場合
		// はスキップする。
		for _, pay := range event.Pays {
			accountsToPay := []model.AccountPay{}
			err := apr.GetAccountPaysByPayId(int(pay.ID), &accountsToPay)
			if err != nil {
				return err
			}
			if len(accountsToPay) == 0 {
				return fmt.Errorf("no accounts to pay for pay ID: %d", pay.ID)
			}
			amountPerPerson := int(math.Round(float64(pay.Amount) / float64(len(accountsToPay))))
			accountsToPayMap := make(map[int]bool)
			for _, accountPay := range accountsToPay {
				accountsToPayMap[int(accountPay.AccountID)] = true
			}
			for _, calc := range calcs {
				if uint(calc.AccountId) == pay.PaidUserID || !accountsToPayMap[calc.AccountId] {
					continue
				}
				calc.AccountAmounts[int(pay.PaidUserID)] += int(amountPerPerson)
			}
		}
		return nil
	}

	err := cu.br.Atomic(atomicBlock)
	if err != nil {
		return nil, err
	}
	return calcs, nil
}
