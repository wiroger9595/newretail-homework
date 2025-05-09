package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"newretail-homework/model"

	"github.com/redis/go-redis/v9"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 從 .env 檔案中取得 DB 設定
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// 設定 PostgreSQL DSN（Data Source Name）
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	// 連接 PostgreSQL 資料庫
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	// 自動遷移資料表
	db.AutoMigrate(&model.Customer{}, &model.Purchase{}, &model.Coupon{}, &model.UserCoupon{})

	return db
}


func InitRedis() *redis.Client {
	// 建立 Redis 連線
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // 若使用 Docker Compose，則可以設定為 "redis:6379"
	})

	// 檢查是否成功連線到 Redis
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	// 返回 Redis client
	return rdb
}