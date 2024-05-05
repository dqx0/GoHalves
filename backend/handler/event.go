package handler

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"
)

type IEventHandler interface {
	GetEventById() gin.HandlerFunc
	GetEventByAccountId() gin.HandlerFunc
	CreateEvent() gin.HandlerFunc
	AddAccountToEvent() gin.HandlerFunc
	UpdateEvent() gin.HandlerFunc
	UpdateAuthority() gin.HandlerFunc
	DeleteEvent() gin.HandlerFunc
	DeleteAccountFromEvent() gin.HandlerFunc

	Calc() gin.HandlerFunc
}
type eventHandler struct {
	bu usecase.IBaseUsecase
}

func NewEventHandler(bu usecase.IBaseUsecase) IEventHandler {
	return &eventHandler{bu}
}
func (ec *eventHandler) GetEventById() gin.HandlerFunc {
	return func(c *gin.Context) {
		var event model.Event
		eu := ec.bu.GetEventUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		event, err = eu.GetEventById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"event": event})
	}
}
func (ec *eventHandler) GetEventByAccountId() gin.HandlerFunc {
	return func(c *gin.Context) {
		var events []model.Event
		eu := ec.bu.GetEventUsecase()
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		events, err = eu.GetEventsByAccountId(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"events": events})
	}
}
func (ec *eventHandler) CreateEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newEvent model.Event
		idInt, ok := getUserIdFromContext(c)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user id"})
			return
		}
		if err := c.ShouldBindJSON(&newEvent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		eu := ec.bu.GetEventUsecase()
		createdEvent, err := eu.CreateEvent(newEvent, idInt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"event": createdEvent})
	}
}
func (ec *eventHandler) AddAccountToEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		accountIdStr := c.Param("account_id")
		eventIdStr := c.Param("event_id")
		authorityIdStr := c.PostForm("authority_id")
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
		authorityId, err := strconv.Atoi(authorityIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authority ID"})
			return
		}
		eu := ec.bu.GetEventUsecase()
		accountEvent, err := eu.AddAccountToEvent(accountId, eventId, authorityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account_event": accountEvent})

	}
}
func (ec *eventHandler) UpdateEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		updateEvent := model.Event{
			ID:          uint(id),
			Title:       c.PostForm("title"),
			Description: c.PostForm("description"),
		}
		eu := ec.bu.GetEventUsecase()
		updatedEvent, err := eu.UpdateEvent(id, updateEvent)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"event": updatedEvent})
	}
}
func (ec *eventHandler) UpdateAuthority() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		authorityId, err := strconv.Atoi(c.PostForm("authority_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid authority ID"})
			return
		}
		eu := ec.bu.GetEventUsecase()
		updatedAuthority, err := eu.UpdateAuthority(id, authorityId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"authority": updatedAuthority})
	}
}
func (ec *eventHandler) DeleteEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		eu := ec.bu.GetEventUsecase()
		deletedEventId, err := eu.DeleteEvent(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Event deleted: " + strconv.Itoa(deletedEventId)})
	}
}
func (ec *eventHandler) DeleteAccountFromEvent() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		accountId, err := strconv.Atoi(c.Param("account_id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
			return
		}
		eu := ec.bu.GetEventUsecase()
		deletedAccountEvent, err := eu.DeleteAccountFromEvent(id, accountId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"account_event": deletedAccountEvent})
		return
	}
}
func (ec *eventHandler) Calc() gin.HandlerFunc {
	return func(c *gin.Context) {
		eventIdStr := c.Param("id")
		eventId, err := strconv.Atoi(eventIdStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
			return
		}
		cu := ec.bu.GetCalcUsecase()
		calcs, err := cu.CalculatePaymentForAccounts(eventId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"calcs": calcs})
		return
	}
}
