package usecase

import (
	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type ISessionUsecase interface {
	Login(username, password string) (bool, error)
}

type sessionUsecase struct {
	br repository.IBaseRepository
}

func NewSessionUsecase(br repository.IBaseRepository) ISessionUsecase {
	return &sessionUsecase{br}
}

func (su *sessionUsecase) Login(username string, password string) (bool, error) {
	sr := su.br.GetAccountRepository()
	user := model.Account{}
	err := sr.GetAccountByUserId(username, &user)
	if err != nil {
		return false, err
	}

	if password == user.Password {
		return true, nil
	}
	// パスワードが一致する場合、trueを返します。
	return false, nil
}
