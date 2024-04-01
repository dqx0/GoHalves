package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"
)

type IAccountHandler interface {
	GetAccountById() gin.HandlerFunc
	CreateAccount() gin.HandlerFunc
	UpdateAccount() gin.HandlerFunc
	DeleteAccount() gin.HandlerFunc
}
type accountHandler struct {
	bu usecase.IBaseUsecase
}

func NewAccountHandler(bu usecase.IBaseUsecase) IAccountHandler {
	return &accountHandler{bu}
}
func (ac *accountHandler) GetAccountById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account model.Account
		au := ac.bu.GetAccountUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		account, err = au.GetAccountById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account": account})
	}
}
func (ac *accountHandler) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account model.Account
		au := ac.bu.GetAccountUsecase()
		err := c.BindJSON(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		createdAccount, err := au.CreateAccount(account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account": createdAccount})
	}
}
func (ac *accountHandler) UpdateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account model.Account
		au := ac.bu.GetAccountUsecase()
		err := c.BindJSON(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		updatedAccount, err := au.UpdateAccount(int(account.ID), account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account": updatedAccount})
	}
}
func (ac *accountHandler) DeleteAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account model.Account
		au := ac.bu.GetAccountUsecase()
		err := c.BindJSON(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		deletedAccount, err := au.DeleteAccount(int(account.ID))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account": deletedAccount})
	}
}
