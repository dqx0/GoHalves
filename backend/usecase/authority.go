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
	br repository.IBaseRepository
}

func NewAuthorityUsecase(br repository.IBaseRepository) IAuthorityUsecase {
	return &authorityUsecase{br}
}
func (au *authorityUsecase) GetAuthorities() ([]model.Authority, error) {
	ar := au.br.GetAuthorityRepository()
	authorities := []model.Authority{}
	if err := ar.GetAuthorities(&authorities); err != nil {
		return nil, err
	}
	return authorities, nil
}
func (au *authorityUsecase) GetAuthorityById(authorityId int) (model.Authority, error) {
	ar := au.br.GetAuthorityRepository()
	authority := model.Authority{}
	if err := ar.GetAuthorityById(authorityId, &authority); err != nil {
		return model.Authority{}, err
	}
	return authority, nil
}
