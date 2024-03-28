package api_v1

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/gin-gonic/gin"
)

type IAccountUsecase interface {
	GetAccounts(accounts []*model.Account) error
	AddAccount(account *model.Account) error
	UpdateAccount(id int, account *model.Account) error
	DeleteAccount(id int, account *model.Account) error
}
type accountUsecase struct {
	accountRepository repository.IAccountRepository
}

func GetAccounts(c *gin.Context) {
	var accounts []*model.Account
	err := repository.GetAccounts(accounts)
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
	savedAccount, err := repository.AddAccount(&account)
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
	savedAccount, err := repository.UpdateAccount(id, &account)
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

	account, err := repository.DeleteAccount(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"account": account})
}
