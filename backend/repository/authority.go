package repository

import (
	"github.com/dqx0/GoHalves/go/model"
	"gorm.io/gorm"
)

type IAuthorityRepository interface {
	GetAuthorities(authorities *[]model.Authority) error
	GetAuthorityById(authorityId int, authority *model.Authority) error
}
type authorityRepository struct {
	db *gorm.DB
}

func NewAuthorityRepository(db *gorm.DB) IAuthorityRepository {
	return &authorityRepository{db}
}
func (aur *authorityRepository) GetAuthorities(authorities *[]model.Authority) error {
	if err := aur.db.Find(authorities).Error; err != nil {
		return err
	}
	return nil
}

func (aur *authorityRepository) GetAuthorityById(authorityId int, authority *model.Authority) error {
	if err := aur.db.Where(&model.Authority{ID: uint(authorityId)}).Find(authority).Error; err != nil {
		return err
	}
	return nil
}
