package api_v1

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/dqx0/GoHalves/go/query"
	"github.com/gin-gonic/gin"
)

func GetEventsByUserId(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events, err := query.GetEventsByUserId(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"events": events})
}
func GetEventById(c *gin.Context) {
	eventIdStr := c.Query("id")
	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := query.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"event": event})
}
func AddEvent(c *gin.Context) {
	var inputEvent *model.InputEvent
	if err := c.BindJSON(&inputEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var event *model.Event
	var accountEvent *model.AccountEvent
	event.Title = inputEvent.Title
	event.Description = inputEvent.Description
	accountEvent.AccountID = uint(inputEvent.UserId)
	accountEvent.AuthorityID = 1
	event, err := query.AddEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	accountEvent, err = query.AddAccountEvent(accountEvent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}
