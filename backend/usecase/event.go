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
	br repository.IBaseRepository
	bv validator.IBaseValidator
}

func NewEventUsecase(br repository.IBaseRepository, bv validator.IBaseValidator) IEventUsecase {
	return &eventUsecase{br, bv}
}
func (eu *eventUsecase) GetEventById(eventId int) (model.Event, error) {
	event := model.Event{}
	er := eu.br.GetEventRepository()
	if err := er.GetEventById(eventId, &event); err != nil {
		return model.Event{}, nil
	}
	return event, nil
}
func (eu *eventUsecase) GetEventsByAccountId(accountId int) ([]model.Event, error) {
	events := []model.Event{}
	er := eu.br.GetEventRepository()
	if err := er.GetEventsByAccountId(accountId, &events); err != nil {
		return nil, err
	}
	return events, nil
}
func (eu *eventUsecase) CreateEvent(event model.Event, createdAccountId int) (model.Event, error) {
	er := eu.br.GetEventRepository()
	if err := er.CreateEvent(&event); err != nil {
		return model.Event{}, err
	}
	accountEvent := model.AccountEvent{
		AccountID:   uint(createdAccountId),
		EventID:     event.ID,
		AuthorityID: 0,
	}
	aer := eu.br.GetAccountEventRepository()
	if err := aer.CreateAccountEvent(&accountEvent); err != nil {
		if err := er.DeleteEvent(int(event.ID), &event); err != nil {
			return model.Event{}, err
		}
		return model.Event{}, err
	}
	return event, nil
}
func (eu *eventUsecase) UpdateEvent(eventId int, event model.Event) (model.Event, error) {
	er := eu.br.GetEventRepository()
	if err := er.UpdateEvent(eventId, &event); err != nil {
		return event, err
	}
	return event, nil
}
func (eu *eventUsecase) DeleteEvent(eventId int, event model.Event) (model.Event, error) {
	// todo: トランザクション
	// 複数のリポジトリを扱う場合はトランザクションを使う
	// トランザクションを使う場合は、トランザクション内でリポジトリを生成する
	// トランザクションはリポジトリのメソッド内で行い、ユースケース層でトランザクションを意識する必要はない
	// 参考: https://gorm.io/ja_JP/docs/transactions.html, https://sano11o1.com/posts/handle-transaction-in-usecase-layer
	// 例　↓
	atomicBlock := func(br repository.IBaseRepository) error {
		er := br.GetEventRepository()
		if err := er.DeleteEvent(eventId, &event); err != nil {
			return err
		}
		return nil
	}
	err := eu.br.Atomic(atomicBlock)
	return event, err
}
