package query

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/gin-gonic/gin"
)

type InputAccountEvent struct {
	AccountID   uint `json:"account_id" binding:"required"`
	EventID     uint `json:"event_id" binding:"required"`
	AuthorityID uint `json:"authority_id" binding:"required"`
}

func GetAccountEvents(c *gin.Context) {
	var accountEvents []model.AccountEvent

	db := gormConnect()

	db.Find(&accountEvents)

	c.JSON(http.StatusOK, gin.H{"account_events": accountEvents})
}

/*
	func AddAccountEvent(c *gin.Context) {
		var inputAccountEvent InputAccountEvent
		if err := c.BindJSON(&inputAccountEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var accountEvent model.AccountEvent
		db := gormConnect()
		if err := db.Create(&accountEvent).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"account_event": accountEvent})
	}
*/
func AddAccountEvent(accountEvent *model.AccountEvent) (*model.AccountEvent, error) {
	db := gormConnect()
	if err := db.Create(&accountEvent).Error; err != nil {
		return nil, err
	}
	return accountEvent, nil
}
func UpdateAccountEvent(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var accountEvent model.AccountEvent

	db.First(&accountEvent, id)

	var inputAccountEvent InputAccountEvent
	c.BindJSON(&inputAccountEvent)

	db.Model(&accountEvent).Update("AccountID", inputAccountEvent.AccountID)
	db.Model(&accountEvent).Update("EventID", inputAccountEvent.EventID)
	db.Model(&accountEvent).Update("AuthorityID", inputAccountEvent.AuthorityID)

	c.JSON(http.StatusOK, gin.H{"account_event": accountEvent})
}
func DeleteAccountEvent(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var accountEvent model.AccountEvent

	db.Delete(&accountEvent, id)

	c.JSON(http.StatusOK, gin.H{"account_event": accountEvent})
}
