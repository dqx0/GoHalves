package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAccountPayRepository interface {
	GetAccountPaysByAccountId(accountId int, accountPays *[]model.AccountPay) error
	GetAccountPaysByPayId(payId int, accountPays *[]model.AccountPay) error
	CreateAccountPay(accountPay *model.AccountPay) error
	UpdateAccountPay(id int, accountPay *model.AccountPay) error
	DeleteAccountPay(id int, accountPay *model.AccountPay) error
}
type accountPayRepository struct {
	db *gorm.DB
}

func NewAccountPayRepository(db *gorm.DB) IAccountPayRepository {
	return &accountPayRepository{db: db}
}
func (apr *accountPayRepository) GetAccountPaysByAccountId(accountId int, accountPays *[]model.AccountPay) error {
	if err := apr.db.Where("account_id = ?", accountId).Find(&accountPays).Error; err != nil {
		return err
	}
	return nil
}
func (apr *accountPayRepository) GetAccountPaysByPayId(payId int, accountPays *[]model.AccountPay) error {
	if err := apr.db.Where("pay_id = ?", payId).Find(&accountPays).Error; err != nil {
		return err
	}
	return nil
}
func (apr *accountPayRepository) CreateAccountPay(accountPay *model.AccountPay) error {
	if err := apr.db.Create(&accountPay).Error; err != nil {
		return err
	}
	return nil
}
func (apr *accountPayRepository) UpdateAccountPay(id int, accountPay *model.AccountPay) error {
	if err := apr.db.Model(&accountPay).Where("id = ?", id).Updates(accountPay).Error; err != nil {
		return err
	}
	return nil
}
func (apr *accountPayRepository) DeleteAccountPay(id int, accountPay *model.AccountPay) error {
	if err := apr.db.Delete(&accountPay, id).Error; err != nil {
		return err
	}
	return nil
}
