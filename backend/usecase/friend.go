package usecase

import (
	"time"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type IFriendUsecase interface {
	GetFriendById(id int) (model.Friend, error)
	GetFriendsByAccountId(accountId int, friends []model.Friend) ([]model.Friend, error)
	SendFriendRequest(sendAccountId int, receivedAccountId int) (model.Friend, error)
	AcceptFriend(id int) (model.Friend, error)
	DeleteFriend(id int) (model.Friend, error)
}
type friendUsecase struct {
	br repository.IBaseRepository
}

func NewFriendUsecase(br repository.IBaseRepository) IFriendUsecase {
	return &friendUsecase{br}
}
func (fu *friendUsecase) GetFriendById(id int) (model.Friend, error) {
	fr := fu.br.GetFriendRepository()
	friend := model.Friend{}
	if err := fr.GetFriendById(id, &friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
func (fu *friendUsecase) GetFriendsByAccountId(accountId int, friends []model.Friend) ([]model.Friend, error) {
	fr := fu.br.GetFriendRepository()
	if err := fr.GetFriendsByAccountId(accountId, &friends); err != nil {
		return nil, err
	}
	return friends, nil
}
func (fu *friendUsecase) SendFriendRequest(sendAccountId int, receivedAccountId int) (model.Friend, error) {
	fr := fu.br.GetFriendRepository()
	var friend model.Friend
	friend.SendAccountID = uint(sendAccountId)
	friend.ReceivedAccountID = uint(receivedAccountId)
	t, err := time.Parse("2006-01-02 15:04:05", "9999-12-31 23:59:59")
	if err != nil {
		return model.Friend{}, err
	}
	friend.AcceptedAt = t
	if err := fr.CreateFriend(&friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
func (fu *friendUsecase) AcceptFriend(id int) (model.Friend, error) {
	fr := fu.br.GetFriendRepository()
	friend := model.Friend{}
	err := fr.GetFriendById(id, &friend)
	if err != nil {
		return model.Friend{}, err
	}
	friend.AcceptedAt = time.Now()
	if err := fr.UpdateFriend(id, &friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
func (fu *friendUsecase) DeleteFriend(id int) (model.Friend, error) {
	fr := fu.br.GetFriendRepository()
	friend := model.Friend{}
	if err := fr.DeleteFriend(id, &friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
