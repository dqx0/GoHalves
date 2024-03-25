package query

import (
	"github.com/dqx0/GoHalves/go/model"
)

func GetAccounts() ([]*model.Account, error) {
	var users []*model.Account

	db := gormConnect()

	// db.Findの結果をチェックして、エラーがあれば返す
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	// usersスライスとnilエラーを返す
	return users, nil
}

func AddAccount(account *model.Account) (*model.Account, error) {
	db := gormConnect()

	if err := db.Create(&account).Error; err != nil {
		return nil, err
	}

	return account, nil
}
func UpdateAccount(id int, account *model.Account) (*model.Account, error) {

	db := gormConnect()

	db.First(&account, id)

	db.Model(&account).Update("UserID", account.UserID)
	db.Model(&account).Update("Name", account.Name)
	db.Model(&account).Update("Email", account.Email)
	db.Model(&account).Update("Password", account.Password)

	return account, nil
}
func DeleteAccount(id int) (*model.Account, error) {
	var account model.Account
	db := gormConnect()

	db.Delete(&account, id)

	return &account, nil
}
