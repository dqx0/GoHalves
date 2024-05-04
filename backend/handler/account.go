package handler

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

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
		id, ok := c.Get("userId")
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user id"})
			return
		}
		idUint, ok := id.(uint)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "User id is not uint"})
			return
		}
		account, err := au.GetAccountById(int(idUint))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		account.Password = ""
		c.JSON(http.StatusOK, gin.H{"account": account})
	}
}

func (ac *accountHandler) CreateAccount() gin.HandlerFunc {
	return func(c *gin.Context) {
		var account model.Account
		au := ac.bu.GetAccountUsecase()

		// JSONからアカウント情報をバインド
		err := c.BindJSON(&account)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		// パスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		account.Password = string(hashedPassword)

		// アカウントを作成
		createdAccount, err := au.CreateAccount(account)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// パスワード情報を除外してレスポンスを返す
		createdAccount.Password = ""
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
		updatedAccount.Password = ""
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
		c.JSON(http.StatusOK, gin.H{"message": deletedAccount.Name + "was deleted"})
	}
}
