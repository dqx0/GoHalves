package router

import (
	"github.com/dqx0/GoHalves/go/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(bh handler.IBaseHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/login", bh.GetSessionHandler().Login())
	r.POST("/account", bh.GetAccountHandler().CreateAccount())
	authorized := r.Group("/")
	authorized.Use(bh.GetSessionHandler().CheckSession)
	{
		// Account
		authorized.GET("/account/:id", bh.GetAccountHandler().GetAccountById())
		authorized.PUT("/account/:id", bh.GetAccountHandler().UpdateAccount())
		authorized.DELETE("/account/:id", bh.GetAccountHandler().DeleteAccount())
		// Event
		authorized.GET("/event/:id", bh.GetEventHandler().GetEventById())
		authorized.GET("/event/account/:id", bh.GetEventHandler().GetEventByAccountId())
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
		authorized.GET("/friend/:id", bh.GetFriendHandler().GetFriendsByAccountId())
		authorized.POST("/friend", bh.GetFriendHandler().SendFriendRequest())
		authorized.PUT("/friend/:id", bh.GetFriendHandler().AcceptFriend())
		authorized.DELETE("/friend/:id", bh.GetFriendHandler().DeleteFriend())
	}
	return r
}
