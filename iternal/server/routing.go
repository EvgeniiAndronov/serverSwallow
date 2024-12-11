package server

import "github.com/gin-gonic/gin"

func SetUpRouts(router *gin.Engine) {

	router.GET("/ws", HandleWebSocket)
	router.GET("/", HandleInfoFromWebSocket)
	router.GET("/clear", HandlerClearData)
}
