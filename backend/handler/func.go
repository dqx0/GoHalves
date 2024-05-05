package handler

import (
	"github.com/gin-gonic/gin"
)

func getUserIdFromContext(c *gin.Context) (int, bool) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, false
	}
	idUint, ok := id.(uint)
	if !ok {
		return 0, false
	}
	idInt := int(idUint)
	return idInt, true
}
