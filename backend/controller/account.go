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
	Login(c *gin.Context) error
	CreateAccount(c *gin.Context) error
	UpdateAccount(c *gin.Context) error
	DeleteAccount(c *gin.Context) error
}
type accountController struct {
	bu usecase.IBaseUsecase
}

func NewAccountController(bu usecase.IBaseUsecase) IAccountController {
	return &accountController{bu}
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
func (ac *accountController) Login(c *gin.Context) error {
	var account model.Account
	au := ac.bu.GetAccountUsecase()
	if err := c.BindJSON(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return err
	}
	account, err := au.Login(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return err
	}
	c.JSON(http.StatusOK, gin.H{"account": account})
	return nil
}
