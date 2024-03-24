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

func AddAccount(inputAccount *model.InputAccount) (*model.Account, error) {
	var account model.Account
	account.UserID = inputAccount.UserID
	account.Name = inputAccount.Name
	account.Email = inputAccount.Email
	account.Password = inputAccount.Password

	db := gormConnect()

	if err := db.Create(&account).Error; err != nil {
		return nil, err
	}

	return &account, nil
}
func UpdateAccount(id int, inputAccount *model.InputAccount) (*model.Account, error) {
	var account model.Account
	account.UserID = inputAccount.UserID
	account.Name = inputAccount.Name
	account.Email = inputAccount.Email
	account.Password = inputAccount.Password

	db := gormConnect()

	db.First(&account, id)

	db.Model(&account).Update("UserID", inputAccount.UserID)
	db.Model(&account).Update("Name", inputAccount.Name)
	db.Model(&account).Update("Email", inputAccount.Email)
	db.Model(&account).Update("Password", inputAccount.Password)

	return &account, nil
}
func DeleteAccount(id int) (*model.Account, error) {
	var account model.Account
	db := gormConnect()

	db.Delete(&account, id)

	return &account, nil
}
