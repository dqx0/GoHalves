package usecase

import (
	"fmt"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"golang.org/x/crypto/bcrypt"
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
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	// パスワードが一致する場合、trueを返します。
	return true, nil
}
