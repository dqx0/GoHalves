package query

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/gin-gonic/gin"
)

type InputAccountPay struct {
	AccountID uint `json:"account_id" binding:"required"`
	PayID     uint `json:"pay_id" binding:"required"`
}

func GetAccountPays(c *gin.Context) {
	var accountPays []model.AccountPay

	db := gormConnect()

	db.Find(&accountPays)

	c.JSON(http.StatusOK, gin.H{"account_pays": accountPays})
}
func AddAccountPay(c *gin.Context) {
	var inputAccountPay InputAccountPay
	if err := c.BindJSON(&inputAccountPay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var accountPay model.AccountPay
	db := gormConnect()
	if err := db.Create(&accountPay).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account_pay": accountPay})
}
func UpdateAccountPay(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var accountPay model.AccountPay

	db.First(&accountPay, id)

	var inputAccountPay InputAccountPay
	c.BindJSON(&inputAccountPay)

	db.Model(&accountPay).Update("AccountID", inputAccountPay.AccountID)
	db.Model(&accountPay).Update("PayID", inputAccountPay.PayID)

	c.JSON(http.StatusOK, gin.H{"account_pay": accountPay})
}
func DeleteAccountPay(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var accountPay model.AccountPay

	db.Delete(&accountPay, id)

	c.JSON(http.StatusOK, gin.H{"account_pay": accountPay})
}
