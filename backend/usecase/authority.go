package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type IAuthorityUsecase interface {
	GetAuthorities() ([]model.Authority, error)
	GetAuthorityById(authorityId int) (model.Authority, error)
}
type authorityUsecase struct {
	ar repository.IAuthorityRepository
}

func NewAuthorityUsecase(ar repository.IAuthorityRepository) IAuthorityUsecase {
	return &authorityUsecase{ar}
}
func (au *authorityUsecase) GetAuthorities() ([]model.Authority, error) {
	authorities := []model.Authority{}
	if err := au.ar.GetAuthorities(&authorities); err != nil {
		return nil, err
	}
	return authorities, nil
}
func (au *authorityUsecase) GetAuthorityById(authorityId int) (model.Authority, error) {
	authority := model.Authority{}
	if err := au.ar.GetAuthorityById(authorityId, &authority); err != nil {
		return model.Authority{}, err
	}
	return authority, nil
}
