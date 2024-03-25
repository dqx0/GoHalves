package api_v1

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/query"
	"github.com/gin-gonic/gin"
)

func GetAccounts(c *gin.Context) {
	accounts, err := query.GetAccounts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}
func AddAccount(c *gin.Context) {
	var inputAccount *model.InputAccount
	if err := c.BindJSON(&inputAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var account model.Account
	account.UserID = inputAccount.UserID
	account.Name = inputAccount.Name
	account.Email = inputAccount.Email
	account.Password = inputAccount.Password
	account.IsBot = false
	savedAccount, err := query.AddAccount(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": savedAccount})
}
func UpdateAccount(c *gin.Context) {
	var inputAccount *model.InputAccount
	if err := c.BindJSON(&inputAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)
	var account model.Account
	account.UserID = inputAccount.UserID
	account.Name = inputAccount.Name
	account.Email = inputAccount.Email
	account.Password = inputAccount.Password
	account.IsBot = false
	savedAccount, err := query.UpdateAccount(id, &account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"account": savedAccount})
}
func DeleteAccount(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	account, err := query.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}
