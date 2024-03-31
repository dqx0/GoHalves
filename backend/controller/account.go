package controller

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"
)

type IAccountController interface {
	GetAccountInfo(c *gin.Context) error
	CreateAccount(c *gin.Context) error
	UpdateAccount(c *gin.Context) error
	DeleteAccount(c *gin.Context) error
}
type accountController struct {
	bu usecase.IBaseUsecase
	su usecase.ISessionUsecase
}

func NewAccountController(bu usecase.IBaseUsecase, su usecase.ISessionUsecase) IAccountController {
	return &accountController{bu, su}
}
func (ac *accountController) GetAccountInfo(c *gin.Context) error {
	var account model.Account
	au := ac.bu.GetAccountUsecase()
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return err
	}
	account, err = au.GetAccountById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"account": account})
	return nil
}
func (ac *accountController) CreateAccount(c *gin.Context) error {
	var account model.Account
	au := ac.bu.GetAccountUsecase()
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return err
	}
	createdAccount, err := au.CreateAccount(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"account": createdAccount})
	return nil
}
func (ac *accountController) UpdateAccount(c *gin.Context) error {
	var account model.Account
	au := ac.bu.GetAccountUsecase()
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return err
	}
	updatedAccount, err := au.UpdateAccount(int(account.ID), account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"account": updatedAccount})
	return nil
}
func (ac *accountController) DeleteAccount(c *gin.Context) error {
	var account model.Account
	au := ac.bu.GetAccountUsecase()
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return err
	}
	deletedAccount, err := au.DeleteAccount(int(account.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"account": deletedAccount})
	return nil
}
