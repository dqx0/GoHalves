package controller

import (
	"github.com/dqx0/GoHalves/go/usecase"
	"github.com/gin-gonic/gin"

	"net/http"
	"os"
)

type ISessionController interface {
	Login(c *gin.Context) error
	Logout(c *gin.Context) error
}
type sessionController struct {
	su usecase.ISessionUsecase
}

func NewSessionController(su usecase.ISessionUsecase) ISessionController {
	return &sessionController{su}
}
func (sc *sessionController) Login(c *gin.Context) error {
	username := c.PostForm("user_id")
	password := c.PostForm("password")
	err := sc.su.Login(c, username, password)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return err
	}
	cookieKey := os.Getenv("LOGIN_USERID_KEY")
	sessionKey, err := sc.su.CreateSession(c, username /*redisValue*/)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return err
	}
	c.SetCookie(cookieKey, sessionKey, 3600, "/", "", true, true)
	c.Redirect(http.StatusFound, "/")
	return nil
}
func (sc *sessionController) Logout(c *gin.Context) error {
	cookieKey := os.Getenv("LOGIN_USERID_KEY")
	sc.su.DeleteSession(c, cookieKey)
	c.Redirect(http.StatusFound, "/")
	return nil
}
