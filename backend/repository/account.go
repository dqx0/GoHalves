package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccounts(accounts *[]model.Account) error
	GetAccountById(id int, account *model.Account) error
	GetAccountByAccountId(id string, account *model.Account) error
	CreateAccount(account *model.Account) error
	UpdateAccount(id int, account *model.Account) error
	DeleteAccount(id int, account *model.Account) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) IAccountRepository {
	return &accountRepository{db}
}

func (ar *accountRepository) GetAccounts(accounts *[]model.Account) error {
	// db.Findの結果をチェックして、エラーがあれば返す
	if err := ar.db.Find(accounts).Error; err != nil {
		return err
	}

	// usersスライスとnilエラーを返す
	return nil
}
func (ar *accountRepository) GetAccountById(id int, account *model.Account) error {
	if err := ar.db.Where(&model.Account{ID: uint(id)}).Find(&account).Error; err != nil {
		return err
	}

	return nil
}

func (ar *accountRepository) GetAccountByAccountId(id string, account *model.Account) error {
	if err := ar.db.Where(&model.Account{UserID: id}).Find(&account).Error; err != nil {
		return err
	}

	return nil
}

func (ar *accountRepository) CreateAccount(account *model.Account) error {
	if err := ar.db.Create(&account).Error; err != nil {
		return err
	}

	return nil
}
func (ar *accountRepository) UpdateAccount(id int, account *model.Account) error {

	result := ar.db.Model(&model.Account{}).Where("id = ?", id).Updates(account)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (ar *accountRepository) DeleteAccount(id int, account *model.Account) error {

	if err := ar.db.Delete(&account, id).Error; err != nil {
		return err
	}

	return nil
}
