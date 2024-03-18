package query

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/gin-gonic/gin"
)

type InputPay struct {
	PaidUserID uint `json:"paid_user_id" binding:"required"`
	EventID    uint `json:"event_id" binding:"required"`
	Amount     uint `json:"amount" binding:"required"`
}

func GetPays(c *gin.Context) {
	var pays []model.Pay

	db := gormConnect()

	db.Find(&pays)

	c.JSON(http.StatusOK, gin.H{"pays": pays})
}
func AddPay(c *gin.Context) {
	var inputPay InputPay
	if err := c.BindJSON(&inputPay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var pay model.Pay
	db := gormConnect()
	if err := db.Create(&pay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"pay": pay})
}
func UpdatePay(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var pay model.Pay

	db.First(&pay, id)

	var inputPay InputPay
	c.BindJSON(&inputPay)

	db.Model(&pay).Update("PaidUserID", inputPay.PaidUserID)
	db.Model(&pay).Update("EventID", inputPay.EventID)
	db.Model(&pay).Update("Amount", inputPay.Amount)

	c.JSON(http.StatusOK, gin.H{"pay": pay})
}
func DeletePay(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var pay model.Pay

	db.Delete(&pay, id)

	c.JSON(http.StatusOK, gin.H{"pay": pay})
}
