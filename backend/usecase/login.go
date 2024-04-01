package usecase

import (
	"context"
	"errors"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type ISessionUsecase interface {
	Login(c context.Context, username, password string) error
	CheckSession(c context.Context, redisKey string) (bool, error)
	CreateSession(c context.Context, value string) (string, error)
	GetSession(c context.Context, redisKey string) (string, error)
	DeleteSession(c context.Context, redisKey string) error
}

type sessionUsecase struct {
	rr repository.IRedisRepository
	br repository.IBaseRepository
}

func NewSessionUsecase(rr repository.IRedisRepository, br repository.IBaseRepository) ISessionUsecase {
	return &sessionUsecase{rr, br}
}

func (u *sessionUsecase) Login(c context.Context, username, password string) error {
	ar := u.br.GetAccountRepository()
	account := model.Account{}
	err := ar.GetAccountByAccountId(username, &account)
	if err != nil {
		return err
	}

	if account.Password != password {
		return errors.New("invalid password")
	}

	return nil
}
func (u *sessionUsecase) CheckSession(c context.Context, redisKey string) (bool, error) {
	_, err := u.rr.GetValue(c, redisKey)
	if err != nil {
		if err == errors.New("SessionKeyが登録されていません。") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
func (u *sessionUsecase) CreateSession(c context.Context, redisValue string) (string, error) {
	sessionKey, err := u.rr.GenerateSessionKey()
	if err != nil {
		return "", err
	}

	err = u.rr.SaveSession(c, sessionKey, redisValue)
	if err != nil {
		return "", err
	}
	return sessionKey, nil
}
func (u *sessionUsecase) GetSession(c context.Context, redisKey string) (string, error) {
	redisValue, err := u.rr.GetValue(c, redisKey)
	if err != nil {
		return "", err
	}
	return redisValue, nil
}
func (u *sessionUsecase) DeleteSession(c context.Context, redisKey string) error {
	err := u.rr.DeleteSession(c, redisKey)
	if err != nil {
		return err
	}
	return nil
}
