package main

import (
	api_v1 "github.com/dqx0/GoHalves/go/api/v1"
	"github.com/dqx0/GoHalves/go/query"
	"github.com/gin-gonic/gin"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	r.GET("/api/accounts", api_v1.GetAccounts)
	r.POST("/api/accounts", api_v1.AddAccount)
	r.PATCH("/api/accounts/:ID", api_v1.UpdateAccount)
	r.DELETE("/api/accounts/:ID", api_v1.DeleteAccount)
	r.GET("/api/events", query.GetEvents)
	r.GET("/api/friends", query.GetFriends)
	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
