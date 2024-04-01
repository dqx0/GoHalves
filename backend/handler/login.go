package handler

import (
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"

	"net/http"
	"os"
)

type ISessionHandler interface {
	Login() gin.HandlerFunc
	Logout() gin.HandlerFunc
	CheckSession(c *gin.Context) bool
}
type sessionHandler struct {
	su usecase.ISessionUsecase
}

func NewSessionHandler(su usecase.ISessionUsecase) ISessionHandler {
	return &sessionHandler{su}
}
func (sc *sessionHandler) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("user_id")
		password := c.PostForm("password")
		err := sc.su.Login(c, username, password)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		cookieKey := os.Getenv("LOGIN_USERID_KEY")
		sessionKey, err := sc.su.CreateSession(c, username /*redisValue*/)
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.SetCookie(cookieKey, sessionKey, 3600, "/", "", true, true)
		c.Redirect(http.StatusFound, "/")
	}
}
func (sc *sessionHandler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookieKey := os.Getenv("LOGIN_USERID_KEY")
		cookieValue, err := c.Cookie(cookieKey)
		if err != nil {
			c.Redirect(http.StatusFound, "/")
			return
		}
		sc.su.DeleteSession(c, cookieValue)
		c.Redirect(http.StatusFound, "/")
	}
}
func (sc *sessionHandler) CheckSession(c *gin.Context) bool {
	cookieKey := os.Getenv("LOGIN_USERID_KEY")
	cookieValue, err := c.Cookie(cookieKey)
	if err != nil {
		c.Redirect(http.StatusFound, "/")
		return false
	}
	isLoggedIn, err := sc.su.CheckSession(c, cookieValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return isLoggedIn
	}
	return isLoggedIn
}
