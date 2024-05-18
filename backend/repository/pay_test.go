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
