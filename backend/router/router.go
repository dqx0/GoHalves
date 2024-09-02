package router

import (
	"net/http"

	"github.com/dqx0/GoHalves/go/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(bh handler.IBaseHandler) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})
	r.POST("/login", bh.GetSessionHandler().Login())
	r.GET("/logout", bh.GetSessionHandler().Logout())
	r.GET("/:id", bh.GetEventHandler().GetCalcData())
	r.POST("/account", func(c *gin.Context) {
		if bh.GetSessionHandler().IsLoggedIn(c) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Already logged in"})
			return
		}
		c.Next()
	}, bh.GetAccountHandler().CreateAccount())

	authorized := r.Group("/")
	authorized.Use(bh.GetSessionHandler().CheckSession)
	{
		//test
		authorized.GET("/", bh.GetAccountHandler().GetAccountById())
		// Account
		authorized.GET("/account", bh.GetAccountHandler().GetAccountById())
		authorized.POST("/accounts", bh.GetAccountHandler().CreateAccount())
		authorized.PUT("/account", bh.GetAccountHandler().UpdateAccount())
		authorized.DELETE("/account", bh.GetAccountHandler().DeleteAccount())
		// Event
		authorized.GET("/event/:id", bh.GetEventHandler().GetEventById())
		authorized.GET("/event/account", bh.GetEventHandler().GetEventByAccountId())
		authorized.GET("/event/calc/:id", bh.GetEventHandler().GetCalcData())
		authorized.POST("/event", bh.GetEventHandler().CreateEvent())
		authorized.PUT("/event", bh.GetEventHandler().UpdateEvent())
		authorized.PUT("/event/authority", bh.GetEventHandler().UpdateAuthority())
		authorized.DELETE("/event/:id", bh.GetEventHandler().DeleteEvent())
		// Pay
		authorized.GET("/pay/event/:eventId", bh.GetPayHandler().GetPaysByEventId())
		authorized.GET("/pay/:id", bh.GetPayHandler().GetPayById())
		authorized.GET("/pay/account/:accountId/event/:eventId", bh.GetPayHandler().GetPaysByAccountIdAndEventId())
		authorized.POST("/pay", bh.GetPayHandler().CreatePay())
		authorized.PUT("/pay/:id", bh.GetPayHandler().UpdatePay())
		authorized.DELETE("/pay/:id", bh.GetPayHandler().DeletePay())
		authorized.POST("/pay/account/:id", bh.GetPayHandler().AddAccountToPay())
		authorized.DELETE("/pay/account/:id", bh.GetPayHandler().DeleteAccountFromPay())
		// Friend
		authorized.GET("/friend", bh.GetFriendHandler().GetFriendsByAccountId())
		authorized.POST("/friend", bh.GetFriendHandler().SendFriendRequest())
		authorized.PUT("/friend/:id", bh.GetFriendHandler().AcceptFriend())
		authorized.DELETE("/friend/:id", bh.GetFriendHandler().DeleteFriend())
	}
	return r
}
