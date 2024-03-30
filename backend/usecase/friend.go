package usecase

import (
	"time"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
)

type IFriendUsecase interface {
	GetFriendsByAccountId(accountId int, friends []model.Friend) ([]model.Friend, error)
	SendFriendRequest(sendAccountId int, receivedAccountId int) (model.Friend, error)
	AcceptFriend(id int, friend model.Friend) (model.Friend, error)
	DeleteFriend(id int) (model.Friend, error)
}
type friendUsecase struct {
	fr repository.IFriendRepository
}

func NewFriendUsecase(fr repository.IFriendRepository) IFriendUsecase {
	return &friendUsecase{fr}
}
func (fu *friendUsecase) GetFriendsByAccountId(accountId int, friends []model.Friend) ([]model.Friend, error) {
	if err := fu.fr.GetFriendsByAccountId(accountId, &friends); err != nil {
		return nil, err
	}
	return friends, nil
}
func (fu *friendUsecase) SendFriendRequest(sendAccountId int, receivedAccountId int) (model.Friend, error) {
	var friend model.Friend
	friend.SendAccountID = uint(sendAccountId)
	friend.ReceivedAccountID = uint(receivedAccountId)
	t, err := time.Parse("2006-01-02 15:04:05", "9999-12-31 23:59:59")
	if err != nil {
		return model.Friend{}, err
	}
	friend.AcceptedAt = t
	if err := fu.fr.CreateFriend(&friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
func (fu *friendUsecase) AcceptFriend(id int, friend model.Friend) (model.Friend, error) {
	friend.AcceptedAt = time.Now()
	if err := fu.fr.UpdateFriend(id, &friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
func (fu *friendUsecase) DeleteFriend(id int) (model.Friend, error) {
	friend := model.Friend{}
	if err := fu.fr.DeleteFriend(id, &friend); err != nil {
		return model.Friend{}, err
	}
	return friend, nil
}
