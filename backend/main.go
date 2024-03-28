package main

import (
	api_v1 "github.com/dqx0/GoHalves/go/api/v1"
	"github.com/dqx0/GoHalves/go/db"
	"github.com/dqx0/GoHalves/go/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	db.NewDB()

	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/accounts", api_v1.GetAccounts)
		api.POST("/accounts", api_v1.AddAccount)
		api.PATCH("/accounts/:ID", api_v1.UpdateAccount)
		api.DELETE("/accounts/:ID", api_v1.DeleteAccount)
		api.GET("/event/:ID", api_v1.GetEventById)
		api.GET("/events/:ID", api_v1.GetEventsByUserId)
		api.GET("/friends", repository.GetFriends)
	}

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
