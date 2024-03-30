package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IPayRepository interface {
	GetPaysByAccountIdAndEventId(accountId int, eventId int, pays *[]model.Pay) error
	GetPaysByEventId(eventId int, pays *[]model.Pay) error
	GetPayById(id int, pay *model.Pay) error
	CreatePay(pay *model.Pay) error
	UpdatePay(id int, pay *model.Pay) error
	DeletePay(id int, pay *model.Pay) error
}
type payRepository struct {
	db *gorm.DB
}

func NewPayRepository(db *gorm.DB) IPayRepository {
	return &payRepository{db}
}

func (pr *payRepository) GetPayById(id int, pay *model.Pay) error {
	if err := pr.db.First(&pay, id).Error; err != nil {
		return err
	}
	return nil
}
func (pr *payRepository) GetPaysByEventId(eventId int, pays *[]model.Pay) error {
	if err := pr.db.Where("event_id = ?", eventId).Find(&pays).Error; err != nil {
		return err
	}
	return nil
}
func (pr *payRepository) GetPaysByAccountIdAndEventId(accountId int, eventId int, pays *[]model.Pay) error {
	if err := pr.db.Joins("JOIN account_pays ON account_pays.pay_id = pays.id").
		Where("account_pays.account_id = ? AND pays.event_id = ?", accountId, eventId).
		Find(&pays).Error; err != nil {
		return err
	}
	return nil
}
func (pr *payRepository) CreatePay(pay *model.Pay) error {
	if err := pr.db.Create(&pay).Error; err != nil {
		return err
	}
	return nil
}
func (pr *payRepository) UpdatePay(id int, pay *model.Pay) error {
	if err := pr.db.Model(&pay).Where("id = ?", id).Updates(pay).Error; err != nil {
		return err
	}
	return nil
}
func (pr *payRepository) DeletePay(id int, pay *model.Pay) error {
	if err := pr.db.Delete(&pay, id).Error; err != nil {
		return err
	}
	return nil
}
