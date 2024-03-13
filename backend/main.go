package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// ルートURL ("/") に対するGETリクエストをハンドル
	r.GET("/", getHelloWorld)

	r.GET("/api/accounts", getAccounts)
	r.POST("/api/accounts", addAccount)
	r.PATCH("/api/accounts/:ID", updateAccount)
	r.DELETE("/api/accounts/:ID", deleteAccount)
	// 8080ポートでサーバーを起動
	r.Run(":8080")
}

func gormConnect() *gorm.DB {
	// 環境変数から接続情報を取得
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := "postgres" // または環境変数から取得
	dbPort := "5432"     // または環境変数から取得

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	db.AutoMigrate(&Account{}, &Event{}, &Pay{}, &AccountEvent{}, &Authority{}, &Friend{}, &AccountPay{})

	return db
}

func getHelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Hello World!!!"})
}

func getAccounts(c *gin.Context) {
	var users []Account

	db := gormConnect()

	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"users": users})
}
func addAccount(c *gin.Context) {
	type InputAccount struct {
		UserID   string `json:"user_id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email"`
		Password string `json:"password" binding:"required"`
	}
	var inputAccount InputAccount
	if err := c.BindJSON(&inputAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	account := Account{
		UserID:   inputAccount.UserID,
		Name:     inputAccount.Name,
		Email:    inputAccount.Email,
		Password: inputAccount.Password,
	}
	db := gormConnect()

	if err := db.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": account})
}
func updateAccount(c *gin.Context) {
	type InputAccount struct {
		UserID   string `json:"user_id" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email"`
		Password string `json:"password" binding:"required"`
	}
	Id := c.Param("ID")
	id, _ := strconv.Atoi(Id)

	db := gormConnect()

	var account Account

	db.First(&account, id)

	var inputAccount InputAccount
	c.BindJSON(&inputAccount)

	db.Model(&account).Update("UserID", inputAccount.UserID)
	db.Model(&account).Update("Name", inputAccount.Name)
	db.Model(&account).Update("Email", inputAccount.Email)
	db.Model(&account).Update("Password", inputAccount.Password)

	c.JSON(http.StatusOK, gin.H{"account": account})
}
func deleteAccount(c *gin.Context) {
	Id := c.Param("ID")
	id, err := strconv.Atoi(Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter ID"})
		return
	}

	db := gormConnect()

	var account Account

	db.Delete(&account, id)

	c.JSON(http.StatusOK, gin.H{"account": account})
}
