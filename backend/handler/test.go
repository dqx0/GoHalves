package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ITestHandler interface {
	Test() gin.HandlerFunc
}
type testHandler struct{}

func NewTestHandler() ITestHandler {
	return &testHandler{}
}
func (th *testHandler) Test() gin.HandlerFunc {
	fmt.Println("test")
	return func(c *gin.Context) {
		userId, ok := c.Get("userId")
		c.JSON(200, gin.H{
			"message": userId,
			"ok":      ok,
		})
	}
}
