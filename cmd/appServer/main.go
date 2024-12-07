package main

import (
	"github.com/gin-gonic/gin"
	"serverSwallow/iternal/server"
)

func main() {
	//server.SetupDatabase()

	router := gin.Default()
	router.GET("/ws", server.HandleWebSocket)
	router.GET("/", server.HandleInfoFromWebSocket)
	router.Run(":8080")
}
