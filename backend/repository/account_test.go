package repository_test

import (
	"testing"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

func TestCheckAccountInfo(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "testuser", Password: "password"}
	err = repo.CreateAccount(&account)
	assert.NoError(t, err)

	ok, err := repo.CheckAccountInfo("testuser", "password")
	assert.NoError(t, err)
	assert.True(t, ok)
}

func TestGetAccounts(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account1 := model.Account{UserID: "user1"}
	account2 := model.Account{UserID: "user2"}
	assert.NoError(t, repo.CreateAccount(&account1))
	assert.NoError(t, repo.CreateAccount(&account2))

	var accounts []model.Account
	assert.NoError(t, repo.GetAccounts(&accounts))
	assert.Len(t, accounts, 2)
}

func TestGetAccountById(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "user3"}
	assert.NoError(t, repo.CreateAccount(&account))

	var retrievedAccount model.Account
	assert.NoError(t, repo.GetAccountById(int(account.ID), &retrievedAccount))
	assert.Equal(t, account.UserID, retrievedAccount.UserID)
}

func TestGetAccountByUserId(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "user4"}
	assert.NoError(t, repo.CreateAccount(&account))

	var retrievedAccount model.Account
	assert.NoError(t, repo.GetAccountByUserId("user4", &retrievedAccount))
	assert.Equal(t, account.UserID, retrievedAccount.UserID)
}

func TestCreateAccount(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "user5"}
	assert.NoError(t, repo.CreateAccount(&account))

	var retrievedAccount model.Account
	assert.NoError(t, db.First(&retrievedAccount, account.ID).Error)
	assert.Equal(t, account.UserID, retrievedAccount.UserID)
}

func TestUpdateAccount(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "user6", Name: "John"}
	assert.NoError(t, repo.CreateAccount(&account))

	update := model.Account{UserID: "user6", Name: "John Doe"}
	assert.NoError(t, repo.UpdateAccount(int(account.ID), &update))

	var updatedAccount model.Account
	assert.NoError(t, db.First(&updatedAccount, account.ID).Error)
	assert.Equal(t, update.Name, updatedAccount.Name)
}

func TestDeleteAccount(t *testing.T) {
	db, err := setupTestDB()
	defer cleanupDB(db)
	assert.NoError(t, err)

	repo := repository.NewAccountRepository(db)

	account := model.Account{UserID: "user7"}
	assert.NoError(t, repo.CreateAccount(&account))

	assert.NoError(t, repo.DeleteAccount(int(account.ID), &account))

	var retrievedAccount model.Account
	err = db.First(&retrievedAccount, account.ID).Error
	assert.Error(t, err)
	assert.Equal(t, gorm.ErrRecordNotFound, err)
}
