package query

import (
	"net/http"
	"strconv"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/gin-gonic/gin"
)

type InputEvent struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func GetEvents(c *gin.Context) {
	var events []model.Event

	db := gormConnect()

	db.Find(&events)

	c.JSON(http.StatusOK, gin.H{"events": events})
}
func AddEvent(c *gin.Context) {
	var inputEvent InputEvent
	if err := c.BindJSON(&inputEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var event model.Event
	db := gormConnect()
	if err := db.Create(&event).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"event": event})
}
func UpdateEvent(c *gin.Context) {
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var event model.Event

	db.First(&event, id)

	var inputEvent InputEvent
	c.BindJSON(&inputEvent)

	db.Model(&event).Update("Title", inputEvent.Title)
	db.Model(&event).Update("Description", inputEvent.Description)

	c.JSON(http.StatusOK, gin.H{"event": event})
}
func DeleteEvent(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var event model.Event

	db.Delete(&event, id)

	c.JSON(http.StatusOK, gin.H{"event": event})
}
