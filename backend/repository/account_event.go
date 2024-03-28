package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAccountEventRepository interface {
	GetAccountEventsByEventId(eventId int, accountEvents *[]model.AccountEvent) error
	GetAccountEventsByAccountId(accountId int, accountEvents *[]model.AccountEvent) error
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
func (aer *accountEventRepository) GetAccountEventsByEventId(eventId int, accountEvents *[]model.AccountEvent) error {
	if err := aer.db.Where("event_id = ?", eventId).Find(&accountEvents).Error; err != nil {
		return err
	}
	return nil
}
func (aer *accountEventRepository) GetAccountEventsByAccountId(accountId int, accountEvents *[]model.AccountEvent) error {
	if err := aer.db.Where("account_id = ?", accountId).Find(&accountEvents).Error; err != nil {
		return err
	}
	return nil
}

func (aer *accountEventRepository) CreateAccountEvent(accountEvent *model.AccountEvent) error {
	if err := aer.db.Create(&accountEvent).Error; err != nil {
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
