package repository

import (
	"gorm.io/gorm"
)

type IBaseRepository interface {
	Atomic(fn func(IBaseRepository) error) error
	BeginTransaction() *gorm.DB
	GetAccountRepository() IAccountRepository
	GetEventRepository() IEventRepository
	GetPayRepository() IPayRepository
	GetAccountEventRepository() IAccountEventRepository
	GetAccountPayRepository() IAccountPayRepository
	GetAuthorityRepository() IAuthorityRepository
	GetFriendRepository() IFriendRepository
}
type baseRepository struct {
	db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) IBaseRepository {
	return &baseRepository{db}
}
func (br *baseRepository) Atomic(fn func(IBaseRepository) error) error {
	return br.db.Transaction(func(tx *gorm.DB) error {
		return fn(NewBaseRepository(tx))
	})
}
func (br *baseRepository) BeginTransaction() *gorm.DB {
	return br.db.Begin()
}
func (br *baseRepository) GetAccountRepository() IAccountRepository {
	return NewAccountRepository(br.db)
}
func (br *baseRepository) GetEventRepository() IEventRepository {
	return NewEventRepository(br.db)
}
func (br *baseRepository) GetPayRepository() IPayRepository {
	return NewPayRepository(br.db)
}
func (br *baseRepository) GetAccountEventRepository() IAccountEventRepository {
	return NewAccountEventRepository(br.db)
}
func (br *baseRepository) GetAccountPayRepository() IAccountPayRepository {
	return NewAccountPayRepository(br.db)
}
func (br *baseRepository) GetAuthorityRepository() IAuthorityRepository {
	return NewAuthorityRepository(br.db)
}
func (br *baseRepository) GetFriendRepository() IFriendRepository {
	return NewFriendRepository(br.db)
}
