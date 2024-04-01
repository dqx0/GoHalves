package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IFriendRepository interface {
	GetFriendById(id int, friend *model.Friend) error
	GetFriendsBySendAccountId(sendAccountId int, friends *[]model.Friend) error
	GetFriendsByReceivedAccountId(sendAccountId int, friends *[]model.Friend) error
	GetFriendsByAccountId(sendAccountId int, friends *[]model.Friend) error
	CreateFriend(friend *model.Friend) error
	UpdateFriend(id int, friend *model.Friend) error
	DeleteFriend(id int, friend *model.Friend) error
}
type friendRepository struct {
	db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) IFriendRepository {
	return &friendRepository{db}
}
func (frr *friendRepository) GetFriendById(id int, friend *model.Friend) error {
	if err := frr.db.Where("id = ?", id).Find(&friend).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) GetFriendsBySendAccountId(sendAccountId int, friends *[]model.Friend) error {
	if err := frr.db.Where("send_account_id = ?", sendAccountId).Find(&friends).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) GetFriendsByReceivedAccountId(receivedAccountId int, friends *[]model.Friend) error {
	if err := frr.db.Where("received_account_id = ?", receivedAccountId).Find(&friends).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) GetFriendsByAccountId(accountId int, friends *[]model.Friend) error {
	if err := frr.db.Where("send_account_id = ? OR received_account_id = ?", accountId, accountId).Find(&friends).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) CreateFriend(friend *model.Friend) error {
	if err := frr.db.Create(&friend).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) UpdateFriend(id int, friend *model.Friend) error {
	if err := frr.db.Model(&friend).Where("id = ?", id).Updates(friend).Error; err != nil {
		return err
	}
	return nil
}
func (frr *friendRepository) DeleteFriend(id int, friend *model.Friend) error {
	if err := frr.db.Delete(&friend, id).Error; err != nil {
		return err
	}
	return nil
}
