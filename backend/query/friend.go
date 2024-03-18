package query

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/gin-gonic/gin"
)

type InputFriend struct {
	SendAccountID     string `json:"send_account_id" binding:"required"`
	ReceivedAccountID string `json:"received_account_id" binding:"required"`
}

func GetFriends(c *gin.Context) {
	var friends []model.Friend

	db := gormConnect()

	db.Preload("SendAccount").Preload("ReceivedAccount").Find(&friends)

	c.JSON(http.StatusOK, gin.H{"friends": friends})
}
func AddFriend(c *gin.Context) {
	var inputFriend InputFriend
	if err := c.BindJSON(&inputFriend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var friend model.Friend
	db := gormConnect()
	if err := db.Create(&friend).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"friend": friend})
}
func UpdateFriend(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var friend model.Friend

	db.First(&friend, id)

	var inputFriend InputFriend
	c.BindJSON(&inputFriend)

	db.Model(&friend).Update("SendAccountID", inputFriend.SendAccountID)
	db.Model(&friend).Update("ReceivedAccountID", inputFriend.ReceivedAccountID)

	c.JSON(http.StatusOK, gin.H{"friend": friend})
}
func DeleteFriend(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var friend model.Friend

	db.Delete(&friend, id)

	c.JSON(http.StatusOK, gin.H{"friend": friend})
}
