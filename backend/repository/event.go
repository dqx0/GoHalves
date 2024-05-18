package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IEventRepository interface {
	GetEventsByAccountId(accountId int, events *[]model.Event) error
	GetEventById(id int, event *model.Event) error
	CreateEvent(event *model.Event) error
	UpdateEvent(id int, event *model.Event) error
	DeleteEvent(id int, event *model.Event) error
}
type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) IEventRepository {
	return &eventRepository{db}
}

func (er *eventRepository) GetEventsByAccountId(accountId int, events *[]model.Event) error {
	if err := er.db.Joins("join accounts_events on accounts_events.event_id = events.id").
		Joins("join accounts on accounts_events.account_id = accounts.id").
		Where("accounts.id = ?", uint(accountId)).
		Find(events).Error; err != nil {
		return err
	}
	return nil
}

func (er *eventRepository) GetEventById(id int, event *model.Event) error {
	if err := er.db.Where(&model.Event{ID: uint(id)}).Find(&event).Error; err != nil {
		return err
	}

	return nil
}

func (er *eventRepository) CreateEvent(event *model.Event) error {
	if err := er.db.Create(&event).Error; err != nil {
		return err
	}
	return nil
}
func (er *eventRepository) UpdateEvent(id int, event *model.Event) error {
	result := er.db.Model(&model.Event{}).Where("id = ?", id).Updates(event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (er *eventRepository) DeleteEvent(id int, event *model.Event) error {
	if err := er.db.Delete(&event, id).Error; err != nil {
		return err
	}

	return nil
}
