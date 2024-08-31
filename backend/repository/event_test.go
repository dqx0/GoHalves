package repository_test

import (
	"testing"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Account{}, &model.Event{}, &model.AccountEvent{}, &model.Pay{}, &model.Authority{}, &model.Friend{}, &model.AccountPay{}); err != nil {
		return nil, err
	}

	return db, nil
}
func cleanupDB(db *gorm.DB) {
	db.Exec("DROP TABLE IF EXISTS accounts")
	db.Exec("DROP TABLE IF EXISTS events")
	db.Exec("DROP TABLE IF EXISTS accounts_events")
	db.Exec("DROP TABLE IF EXISTS pays")
	db.Exec("DROP TABLE IF EXISTS authorities")
	db.Exec("DROP TABLE IF EXISTS friends")
	db.Exec("DROP TABLE IF EXISTS accounts_pays")
	db.AutoMigrate(&model.Account{}, &model.Event{}, &model.AccountEvent{})
}
func TestGetEventsByAccountId(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)
	repo := repository.NewEventRepository(db)

	// テストデータの挿入
	account := model.Account{UserID: "testuser"}
	db.Create(&account)

	event := model.Event{Title: "Test Event", Description: "This is a test event"}
	db.Create(&event)

	accountEvent := model.AccountEvent{AccountID: account.ID, EventID: event.ID}
	db.Create(&accountEvent)

	// 実際のテスト
	var events []model.Event
	err = repo.GetEventsByAccountId(int(account.ID), &events)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(events))
	assert.Equal(t, "Test Event", events[0].Title)
}
func TestGetEventById(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewEventRepository(db)

	// テストデータの挿入
	event := model.Event{Title: "Test Event", Description: "This is a test event"}
	db.Create(&event)

	// 実際のテスト
	var retrievedEvent model.Event
	err = repo.GetEventById(int(event.ID), &retrievedEvent)
	assert.NoError(t, err)
	assert.Equal(t, event.Title, retrievedEvent.Title)
}
func TestCreateEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewEventRepository(db)

	// 実際のテスト
	event := model.Event{Title: "Test Event", Description: "This is a test event"}
	err = repo.CreateEvent(&event)
	assert.NoError(t, err)

	// データベースからイベントを取得して検証
	var retrievedEvent model.Event
	db.First(&retrievedEvent, event.ID)
	assert.Equal(t, event.Title, retrievedEvent.Title)
}
func TestUpdateEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewEventRepository(db)

	// テストデータの挿入
	event := model.Event{Title: "Old Event", Description: "This is an old event"}
	db.Create(&event)

	// 実際のテスト
	updatedEvent := model.Event{Title: "Updated Event", Description: "This is an updated event"}
	err = repo.UpdateEvent(int(event.ID), &updatedEvent)
	assert.NoError(t, err)

	// データベースからイベントを取得して検証
	var retrievedEvent model.Event
	db.First(&retrievedEvent, event.ID)
	assert.Equal(t, "Updated Event", retrievedEvent.Title)
}
func TestDeleteEvent(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Failed to set up test DB: %v", err)
	}
	defer cleanupDB(db)

	repo := repository.NewEventRepository(db)

	// テストデータの挿入
	event := model.Event{Title: "Test Event", Description: "This is a test event"}
	db.Create(&event)

	// 実際のテスト
	err = repo.DeleteEvent(int(event.ID), &event)
	assert.NoError(t, err)

	// データベースからイベントを取得して検証
	var retrievedEvent model.Event
	result := db.First(&retrievedEvent, event.ID)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}
