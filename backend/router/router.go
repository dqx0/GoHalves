package router

import (
	"net/http"

	"github.com/dqx0/GoHalves/go/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(bc handler.IBaseHandler, sc handler.ISessionHandler) *gin.Engine {
	r := gin.Default()
	r.Use(CheckLoggedIn(sc))
	// Account
	r.GET("/account/:id", bc.GetAccountHandler().GetAccountById())
	r.POST("/account", bc.GetAccountHandler().CreateAccount())
	r.PUT("/account/:id", bc.GetAccountHandler().UpdateAccount())
	r.DELETE("/account/:id", bc.GetAccountHandler().DeleteAccount())
	// Event
	r.GET("/event/:id", bc.GetEventHandler().GetEventById())
	r.GET("/event/account/:id", bc.GetEventHandler().GetEventByAccountId())
	r.POST("/event", bc.GetEventHandler().CreateEvent())
	r.PUT("/event", bc.GetEventHandler().UpdateEvent())
	r.PUT("/event/authority", bc.GetEventHandler().UpdateAuthority())
	r.DELETE("/event/:id", bc.GetEventHandler().DeleteEvent())
	// Pay
	r.GET("/pay/event/:eventId", bc.GetPayHandler().GetPaysByEventId())
	r.GET("/pay/:id", bc.GetPayHandler().GetPayById())
	r.GET("/pay/account/:accountId/event/:eventId", bc.GetPayHandler().GetPaysByAccountIdAndEventId())
	r.POST("/pay", bc.GetPayHandler().CreatePay())
	r.PUT("/pay/:id", bc.GetPayHandler().UpdatePay())
	r.DELETE("/pay/:id", bc.GetPayHandler().DeletePay())
	r.POST("/pay/account/:id", bc.GetPayHandler().AddAccountToPay())
	r.DELETE("/pay/account/:id", bc.GetPayHandler().DeleteAccountFromPay())
	// Friend
	r.GET("/friend/:id", bc.GetFriendHandler().GetFriendsByAccountId())
	r.POST("/friend", bc.GetFriendHandler().SendFriendRequest())
	r.PUT("/friend/:id", bc.GetFriendHandler().AcceptFriend())
	r.DELETE("/friend/:id", bc.GetFriendHandler().DeleteFriend())
	return r
}
func CheckLoggedIn(sc handler.ISessionHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		// セッションIDが存在する場合、そのセッションが有効かどうかを確認
		// ここでは、データベースまたはメモリ内のセッションストアを確認することを想定しています
		if sc.CheckSession(c) {
			// セッションが無効な場合、ログインページにリダイレクト
			c.Redirect(http.StatusTemporaryRedirect, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
