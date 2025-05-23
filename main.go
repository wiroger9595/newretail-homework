package main

import (
	"log"
	config "newretail-homework/config"
	route "newretail-homework/routes"
	"os"

	"github.com/gin-gonic/gin"
)


func main() {

    db := config.InitDB()
	rdb := config.InitRedis() 
	r := gin.Default()
	route.Customer(r, db)
	route.Coupon(r, db, rdb)

    port := os.Getenv("PORT")
    if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	r.Run(":" + port)
	
}