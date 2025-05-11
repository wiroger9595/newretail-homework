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

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	db.AutoMigrate(&model.Customer{}, &model.Purchase{}, &model.Coupon{}, &model.UserCoupon{})

	return db
}


func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}
	return rdb
}