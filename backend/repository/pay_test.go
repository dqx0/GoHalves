package repository_test

import (
	"testing"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetPayById(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewPayRepository(db)

	// テストデータの挿入
	pay := model.Pay{Amount: 1000, EventID: 1}
	db.Create(&pay)

	// 実際のテスト
	var retrievedPay model.Pay
	err = repo.GetPayById(int(pay.ID), &retrievedPay)
	assert.NoError(t, err)
	assert.Equal(t, pay.Amount, retrievedPay.Amount)
	assert.Equal(t, pay.EventID, retrievedPay.EventID)
}
func TestGetPaysByEventId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewPayRepository(db)

	// テストデータの挿入
	pay1 := model.Pay{Amount: 1000, EventID: 1}
	pay2 := model.Pay{Amount: 2000, EventID: 1}
	db.Create(&pay1)
	db.Create(&pay2)

	// 実際のテスト
	var pays []model.Pay
	err = repo.GetPaysByEventId(int(pay1.EventID), &pays)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(pays))
	assert.Equal(t, uint(1000), pays[0].Amount)
	assert.Equal(t, uint(2000), pays[1].Amount)
}
func TestCreatePay(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewPayRepository(db)

	// テストデータの挿入
	pay := model.Pay{Amount: 1000, EventID: 1}

	// 実際のテスト
	err = repo.CreatePay(&pay)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, pay.ID)
}
func TestUpdatePay(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewPayRepository(db)

	// テストデータの挿入
	pay := model.Pay{Amount: 1000, EventID: 1}
	db.Create(&pay)

	// 実際のテスト
	pay.Amount = 2000
	err = repo.UpdatePay(int(pay.ID), &pay)
	assert.NoError(t, err)

	var retrievedPay model.Pay
	db.First(&retrievedPay, pay.ID)
	assert.Equal(t, uint(2000), retrievedPay.Amount)
}
func TestDeletePay(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewPayRepository(db)

	// テストデータの挿入
	pay := model.Pay{Amount: 1000, EventID: 1}
	db.Create(&pay)

	// 実際のテスト
	err = repo.DeletePay(int(pay.ID), &pay)
	assert.NoError(t, err)

	var retrievedPay model.Pay
	err = db.First(&retrievedPay, pay.ID).Error
	assert.Error(t, err)
}
