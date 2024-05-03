package usecase

import (
	"context"

	"github.com/dqx0/GoHalves/go/repository"
)

type ISessionUsecase interface {
	Login(c context.Context, username, password string) (bool, error)
}

type sessionUsecase struct {
	br repository.IBaseRepository
}

func NewSessionUsecase(br repository.IBaseRepository) ISessionUsecase {
	return &sessionUsecase{br}
}

func (u *sessionUsecase) Login(c context.Context, username, password string) (bool, error) {
	ar := u.br.GetAccountRepository()
	ok, err := ar.CheckAccountInfo(username, password)
	if err != nil {
		return false, err
	}

	return ok, nil
}
