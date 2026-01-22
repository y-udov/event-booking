package main

import (
	"github.com/gin-gonic/gin"
	"github.com/y-udov/event-booking/db"
	"github.com/y-udov/event-booking/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
