package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"
)

type IFriendHandler interface {
	GetFriendsByAccountId() gin.HandlerFunc
	SendFriendRequest() gin.HandlerFunc
	AcceptFriend() gin.HandlerFunc
	DeleteFriend() gin.HandlerFunc
}
type friendHandler struct {
	bu usecase.IBaseUsecase
}

func NewFriendHandler(bu usecase.IBaseUsecase) IFriendHandler {
	return &friendHandler{bu}
}
func (fc *friendHandler) GetFriendsByAccountId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var friends []model.Friend
		fu := fc.bu.GetFriendUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		friends, err = fu.GetFriendsByAccountId(id, friends)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"friends": friends})
	}
}
func (fc *friendHandler) SendFriendRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var friend model.Friend
		fu := fc.bu.GetFriendUsecase()
		sendAccountId, err := strconv.Atoi(c.PostForm("sendAccountId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid send account ID"})
			return
		}
		receivedAccountId, err := strconv.Atoi(c.PostForm("receivedAccountId"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid received account ID"})
			return
		}
		friend, err = fu.SendFriendRequest(sendAccountId, receivedAccountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"friend": friend})
	}
}
func (fc *friendHandler) AcceptFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var friend model.Friend
		fu := fc.bu.GetFriendUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
			return
		}
		friend, err = fu.AcceptFriend(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"friend": friend})
	}
}
func (fc *friendHandler) DeleteFriend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var friend model.Friend
		fu := fc.bu.GetFriendUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid friend ID"})
			return
		}
		friend, err = fu.DeleteFriend(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"friend": friend})
	}
}
