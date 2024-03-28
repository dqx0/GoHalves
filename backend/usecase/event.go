package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/dqx0/GoHalves/go/validator"
)

type IEventUsecase interface {
	GetEventById(eventId int) (model.Event, error)
	GetEventsByAccountId(accountId int) ([]model.Event, error)
	CreateEvent(event model.Event, createdAccountId int) (model.Event, error)
	UpdateEvent(eventId int, event model.Event) (model.Event, error)
	DeleteEvent(eventId int, event model.Event) (model.Event, error)
}
type eventUsecase struct {
	er  repository.IEventRepository
	aer repository.IAccountEventRepository
	pr  repository.IPayRepository
	ev  validator.IEventValidator
}

func NewEventUsecase(er repository.IEventRepository, ev validator.IEventValidator) IEventUsecase {
	return &eventUsecase{er, ev}
}
func (eu *eventUsecase) GetEventById(eventId int) (model.Event, error) {
	event := model.Event{}
	if err := eu.er.GetEventById(eventId, &event); err != nil {
		return model.Event{}, nil
	}
	return event, nil
}
func (eu *eventUsecase) GetEventsByAccountId(accountId int) ([]model.Event, error) {
	events := []model.Event{}
	if err := eu.er.GetEventsByAccountId(accountId, &events); err != nil {
		return nil, err
	}
	return events, nil
}
func (eu *eventUsecase) CreateEvent(event model.Event, createdAccountId int) (model.Event, error) {
	if err := eu.er.CreateEvent(&event); err != nil {
		return model.Event{}, err
	}
	accountEvent := model.AccountEvent{
		AccountID:   uint(createdAccountId),
		EventID:     event.ID,
		AuthorityID: 0,
	}
	if err := eu.aer.CreateAccountEvent(&accountEvent); err != nil {
		if err := eu.er.DeleteEvent(int(event.ID), &event); err != nil {
			return model.Event{}, err
		}
		return model.Event{}, err
	}
	return event, nil
}
func (eu *eventUsecase) UpdateEvent(eventId int, event model.Event) (model.Event, error) {
	if err := eu.er.UpdateEvent(eventId, &event); err != nil {
		return event, err
	}
	return event, nil
}
func (eu *eventUsecase) DeleteEvent(eventId int, event model.Event) (model.Event, error) {
	if err := eu.pr.DeletePays(eventId); err != nil {
		return event, err
	}
	accountEvents := []model.AccountEvent{}
	if err := eu.aer.GetAccountEventsByEventId(eventId, &accountEvents); err != nil {
		return
	}
}
