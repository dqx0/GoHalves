package usecase

import (
	"fmt"
	"math"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type ICalcUsecase interface {
	CalculatePaymentForAccounts(eventId int) (model.Calc, error)
}
type calcUsecase struct {
	br repository.IBaseRepository
}

func NewCalcUsecase(br repository.IBaseRepository) ICalcUsecase {
	return &calcUsecase{br}
}
func (cu *calcUsecase) CalculatePaymentForAccounts(eventId int) (model.Calc, error) {
	event := model.Event{}
	er := cu.br.GetEventRepository()
	if err := er.GetEventById(eventId, &event); err != nil {
		return model.Calc{}, err
	}
	calcs := model.Calc{
		EventId:        int(event.ID),
		AccountAmounts: []model.AccountAmount{},
	}
	atomicBlock := func(br repository.IBaseRepository) error {
		pr := br.GetPayRepository()
		aer := br.GetAccountEventRepository()

		// 支払い情報を取得
		if err := pr.GetPaysByEventId(int(event.ID), &event.Pays); err != nil {
			return err
		}

		// イベントに参加しているアカウントを取得
		if err := aer.GetAccountsByEventId(int(event.ID), &event.Accounts); err != nil {
			return err
		}

		// 互いの支払い金額を計算するための構造体を作成
		calcs = cu.initializeCalc(calcs, event)

		// 支払いごとの対象者をaccounts_payから、立替者をpay.PaidUserIDから、
		// 金額をpay.Amountから取得し、金額を対象者の数で割って小数第1位で四捨五入
		// する。AccountIdが対象者のID、AccountAmountsのキーが立替者のID、値が金額
		// となるようにAccountAmountsを更新する。AccountIdが立替者のIDと一致する場合
		// はスキップする。
		calcs, err := cu.calc(calcs, event)
		if err != nil {
			return err
		}

		cu.subtractGap(calcs.AccountAmounts)
		return nil
	}

	err := cu.br.Atomic(atomicBlock)
	if err != nil {
		return model.Calc{}, err
	}
	return calcs, nil
}
func (cu *calcUsecase) getIndexOfTheAccount(accountAmounts []model.AccountAmount, accountId int) int {
	for i, accountAmount := range accountAmounts {
		if accountAmount.AccountId == accountId {
			return i
		}
	}
	return -1
}
func (cu *calcUsecase) initializeCalc(calcs model.Calc, event model.Event) model.Calc {
	for _, account := range event.Accounts {
		calcs.AccountAmounts = append(calcs.AccountAmounts, model.AccountAmount{
			AccountId: int(account.ID),
			Amount:    make(map[int]int),
		})
		for _, otherAccount := range event.Accounts {
			if otherAccount.ID != account.ID {
				i := cu.getIndexOfTheAccount(calcs.AccountAmounts, int(account.ID))
				calcs.AccountAmounts[i].Amount[int(otherAccount.ID)] = 0
			}
		}
	}
	return calcs
}
func (cu *calcUsecase) calc(calcs model.Calc, event model.Event) (model.Calc, error) {
	apr := cu.br.GetAccountPayRepository()
	for _, pay := range event.Pays {
		accountsToPay := []model.AccountPay{}
		err := apr.GetAccountPaysByPayId(int(pay.ID), &accountsToPay)
		if err != nil {
			return model.Calc{}, err
		}
		if len(accountsToPay) == 0 {
			return model.Calc{}, fmt.Errorf("no accounts to pay for pay ID: %d", pay.ID)
		}
		amountPerPerson := int(math.Round(float64(pay.Amount) / float64(len(accountsToPay))))
		accountsToPayMap := make(map[int]bool)
		for _, accountPay := range accountsToPay {
			accountsToPayMap[int(accountPay.AccountID)] = true
		}
		for _, calc := range calcs.AccountAmounts {
			if uint(calc.AccountId) == pay.PaidUserID || !accountsToPayMap[calc.AccountId] {
				continue
			}
			calc.Amount[int(pay.PaidUserID)] += int(amountPerPerson)
		}
	}
	return calcs, nil
}
func (cu *calcUsecase) subtractGap(accountAmounts []model.AccountAmount) []model.AccountAmount {
	for _, calc := range accountAmounts {
		for accountId, amount := range calc.Amount {
			i := cu.getIndexOfTheAccount(accountAmounts, accountId)
			if intMin(amount, accountAmounts[i].Amount[int(calc.AccountId)]) == 0 {
				continue
			}
			gap := intAbs(amount - accountAmounts[i].Amount[int(calc.AccountId)])
			calc.Amount[accountId] -= gap
			accountAmounts[i].Amount[int(calc.AccountId)] -= gap
		}
	}
	return accountAmounts
}

func intMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
