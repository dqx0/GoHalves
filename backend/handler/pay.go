package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"
)

type IPayHandler interface {
	GetPaysByEventId() gin.HandlerFunc
	GetPayById() gin.HandlerFunc
	GetPaysByAccountIdAndEventId() gin.HandlerFunc
	AddAccountToPay() gin.HandlerFunc
	CreatePay() gin.HandlerFunc
	UpdatePay() gin.HandlerFunc
	DeletePay() gin.HandlerFunc
	DeleteAccountFromPay() gin.HandlerFunc
}
type payHandler struct {
	bu usecase.IBaseUsecase
}

func NewPayHandler(bu usecase.IBaseUsecase) IPayHandler {
	return &payHandler{bu}
}
func (pc *payHandler) GetPaysByEventId() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventIdStr := c.Param("eventId")
		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		pu := pc.bu.GetPayUsecase()
		pays, err := pu.GetPaysByEventId(eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pays": pays})
	}
}
func (pc *payHandler) GetPayById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay model.Pay
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pay ID"})
			return
		}
		pay, err = pu.GetPayById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pay": pay})
	}
}
func (pc *payHandler) GetPaysByAccountIdAndEventId() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountIdStr := c.Param("accountId")
		eventIdStr := c.Param("eventId")
		accountId, err := strconv.Atoi(accountIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		pu := pc.bu.GetPayUsecase()
		pays, err := pu.GetPaysByAccountIdAndEventId(accountId, eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pays": pays})
	}
}
func (pc *payHandler) AddAccountToPay() gin.HandlerFunc {
	return func(c *gin.Context) {
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pay ID"})
			return
		}
		accountIdStr := c.Param("account_d")
		accountId, err := strconv.Atoi(accountIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		accountPay, err := pu.AddAccountToPay(id, accountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"accountPay": accountPay})
	}
}
func (pc *payHandler) CreatePay() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay model.Pay
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		accountId, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		if err := c.ShouldBindJSON(&pay); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		accountIdsToPay := []int{}
		if err := c.ShouldBindJSON(&accountIdsToPay); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pay, err = pu.CreatePay(pay, accountId, accountIdsToPay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"pay": pay})
	}
}
func (pc *payHandler) UpdatePay() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay model.Pay
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pay ID"})
			return
		}
		if err := c.ShouldBindJSON(&pay); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		pay, err = pu.UpdatePay(id, pay)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pay": pay})
	}
}
func (pc *payHandler) DeletePay() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pay model.Pay
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pay ID"})
			return
		}
		pay, err = pu.DeletePay(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pay": pay})
	}
}
func (pc *payHandler) DeleteAccountFromPay() gin.HandlerFunc {
	return func(c *gin.Context) {
		pu := pc.bu.GetPayUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pay ID"})
			return
		}
		accountIdStr := c.Param("account_id")
		accountId, err := strconv.Atoi(accountIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		accountPay, err := pu.DeleteAccountFromPay(id, accountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"accountPay": accountPay})
	}
}
