package main

import (
	"newretail-homework/config"
	"newretail-homework/routes"

	"github.com/gin-gonic/gin"
)


func main() {


    r := gin.Default()

	config.ConnectDB()
	routes.RegisterRoutes(r)

	r.Run(":8080")

	
}