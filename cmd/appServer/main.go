package main

import (
	"github.com/gin-gonic/gin"
	"serverSwallow/iternal/server"
)

func main() {
	router := gin.Default()
	server.SetUpRouts(router)
	router.Run(":8080")
}
