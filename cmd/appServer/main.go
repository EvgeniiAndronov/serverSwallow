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
	router.GET("/clear", server.HandlerClearData)

	router.Run(":8080")
}
