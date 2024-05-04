package repository

import (
	"fmt"

	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAccountRepository interface {
	GetAccounts(accounts *[]model.Account) error
	GetAccountById(id int, account *model.Account) error
	GetAccountByUserId(id string, account *model.Account) error
	CheckAccountInfo(userId string, password string) (bool, error)
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

func (ar *accountRepository) CheckAccountInfo(userId string, password string) (bool, error) {
	var account model.Account
	if err := ar.db.Where("user_id = ?", userId).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // ユーザーが見つからない場合はfalseを返す
		}
		return false, err // データベースエラーが発生した場合
	}
	if account.Password == password {
		return true, nil // パスワードが一致する場合はtrueを返す
	}

	return false, nil // パスワードが一致しない場合はfalseを返す
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

func (ar *accountRepository) GetAccountByUserId(id string, account *model.Account) error {
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
	updateData := map[string]interface{}{
		"UserID": account.UserID,
		"Name":   account.Name,
		"Email":  account.Email,
	}

	result := ar.db.Model(&model.Account{}).Where("id = ?", id).Updates(updateData)
	if result.Error != nil {
		fmt.Println(result.Error)
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
