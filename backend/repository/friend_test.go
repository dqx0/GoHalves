package repository_test

import (
	"testing"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/stretchr/testify/assert"
)

func TestGetFriendsByID(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)
	repo := repository.NewFriendRepository(db)

	// テストデータの挿入
	friend := model.Friend{SendAccountID: 1, ReceivedAccountID: 2}
	db.Create(&friend)

	// 実際のテスト
	var retrievedFriend model.Friend
	err = repo.GetFriendById(int(friend.ID), &retrievedFriend)
	assert.NoError(t, err)
	assert.Equal(t, friend.SendAccountID, retrievedFriend.SendAccountID)
	assert.Equal(t, friend.ReceivedAccountID, retrievedFriend.ReceivedAccountID)
}
func TestCreateFriend(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)
	repo := repository.NewFriendRepository(db)

	// テストデータの挿入
	friend := model.Friend{SendAccountID: 1, ReceivedAccountID: 2}

	// 実際のテスト
	err = repo.CreateFriend(&friend)
	assert.NoError(t, err)

	// データベースから友達を取得して検証
	var retrievedFriend model.Friend
	db.First(&retrievedFriend, friend.ID)
	assert.Equal(t, friend.SendAccountID, retrievedFriend.SendAccountID)
	assert.Equal(t, friend.ReceivedAccountID, retrievedFriend.ReceivedAccountID)
}
func TestUpdateFriend(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)
	repo := repository.NewFriendRepository(db)

	// テストデータの挿入
	friend := model.Friend{SendAccountID: 1, ReceivedAccountID: 2}
	db.Create(&friend)

	// 実際のテスト
	friend.SendAccountID = 3
	err = repo.UpdateFriend(int(friend.ID), &friend)
	assert.NoError(t, err)

	var retrievedFriend model.Friend
	db.First(&retrievedFriend, friend.ID)
	assert.Equal(t, uint(3), retrievedFriend.SendAccountID)
}

func TestDeleteFriend(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)
	repo := repository.NewFriendRepository(db)

	// テストデータの挿入
	friend := model.Friend{SendAccountID: 1, ReceivedAccountID: 2}
	db.Create(&friend)

	// 実際のテスト
	err = repo.DeleteFriend(int(friend.ID), &friend)
	assert.NoError(t, err)

	var retrievedFriend model.Friend
	db.First(&retrievedFriend, friend.ID)
	assert.Equal(t, uint(0), retrievedFriend.ID)
}
