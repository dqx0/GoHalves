package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAccountEventRepository interface {
	GetAccountsByEventId(eventId int, accountEvents *[]model.Account) error
	GetEventsByAccountId(accountId int, accountEvents *[]model.Event) error
	GetAccountEventByEventIdAndAccountId(eventId int, accountId int, accountEvent *model.AccountEvent) error
	CreateAccountEvent(accountEvent *model.AccountEvent) error
	UpdateAccountEvent(id int, accountEvent *model.AccountEvent) error
	DeleteAccountEvent(id int, accountEvent *model.AccountEvent) error
}
type accountEventRepository struct {
	db *gorm.DB
}

func NewAccountEventRepository(db *gorm.DB) IAccountEventRepository {
	return &accountEventRepository{db: db}
}
func (aer *accountEventRepository) GetAccountsByEventId(eventId int, accounts *[]model.Account) error {
	if err := aer.db.Table("accounts_events").
		Select("accounts.id, accounts.name, accounts.email, accounts.created_at, accounts.updated_at").
		Joins("JOIN accounts ON accounts_events.account_id = accounts.id").
		Where("accounts_events.event_id = ?", eventId).
		Find(&accounts).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) GetEventsByAccountId(accountId int, events *[]model.Event) error {
	if err := aer.db.Table("accounts_events").
		Select("events.id, events.name, events.description, events.created_at, events.updated_at").
		Joins("JOIN events ON accounts_events.event_id = events.id").
		Where("accounts_events.account_id = ?", accountId).
		Find(&events).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) GetAccountEventByEventIdAndAccountId(eventId int, accountId int, accountEvent *model.AccountEvent) error {
	if err := aer.db.Where("event_id = ? AND account_id = ?", eventId, accountId).Find(&accountEvent).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) CreateAccountEvent(accountEvent *model.AccountEvent) error {
	if err := aer.db.Create(accountEvent).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) UpdateAccountEvent(id int, accountEvent *model.AccountEvent) error {
	if err := aer.db.Model(&accountEvent).Where("id = ?", id).Updates(accountEvent).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) DeleteAccountEvent(id int, accountEvent *model.AccountEvent) error {
	if err := aer.db.Delete(&accountEvent, id).Error; err != nil {
		return err
	}

	return nil
}
