package query

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dqx0/GoHalves/go/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB  // データベース接続のシングルトンインスタンス
	once sync.Once // インスタンスの初期化を一度だけ行うための同期メカニズム
)

func gormConnect() *gorm.DB {
	once.Do(func() {
		// 環境変数から接続情報を取得
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("環境変数の読み込みに失敗しました: %v", err)
		}

		dbUser := os.Getenv("POSTGRES_USER")
		dbPassword := os.Getenv("POSTGRES_PASSWORD")
		dbName := os.Getenv("POSTGRES_DB")
		dbHost := os.Getenv("POSTGRES_HOST")
		dbPort := os.Getenv("POSTGRES_PORT")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", dbHost, dbUser, dbPassword, dbName, dbPort)
		var errOpen error
		db, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if errOpen != nil {
			log.Fatalf("データベース接続に失敗しました: %v", errOpen)
		}

		// モデルの自動マイグレーション
		errAutoMigrate := db.AutoMigrate(&model.Account{}, &model.Event{}, &model.Pay{}, &model.AccountEvent{}, &model.Authority{}, &model.Friend{}, &model.AccountPay{})
		if errAutoMigrate != nil {
			log.Fatalf("自動マイグレーションに失敗しました: %v", errAutoMigrate)
		}
	})

	return db
}
