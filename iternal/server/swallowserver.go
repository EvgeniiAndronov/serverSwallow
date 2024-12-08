package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"serverSwallow/iternal/models"
	"slices"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	onServerData = make([]models.UserData, 0)
)

func HandleInfoFromWebSocket(c *gin.Context) {
	c.JSON(
		200,
		onServerData,
	)
}

func HandlerClearData(c *gin.Context) {

	onServerData = onServerData[:0]
	c.JSON(200, gin.H{"message": "cleared"})
}

func HandleWebSocket(c *gin.Context) {

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not upgrade to WebSocket"})
		return
	}
	defer conn.Close()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading message"})
			break
		}

		var message models.UserData

		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing message"})
			break
		}

		ids := make([]string, 0)
		bulletCoord := make([]models.Coordinates, 0)
		tankCoord := make([]models.Coordinates, 0)

		if len(onServerData) == 0 {
			onServerData = append(onServerData, message)
		}

		for i := range onServerData {
			fmt.Printf("\n---%+v---%+v---\n", onServerData[i].IdUser, message.IdUser)
			if onServerData[i].IdUser == message.IdUser {
				onServerData[i].IsAlive = message.IsAlive
				onServerData[i].BulletX = message.BulletX
				onServerData[i].BulletY = message.BulletY
				onServerData[i].TankX = message.TankX
				onServerData[i].TankY = message.TankY
				onServerData[i].AnimateStatus = message.AnimateStatus
				onServerData[i].AngleTank = message.AngleTank
			}
			ids = append(ids, onServerData[i].IdUser)
			bulletCoord = append(bulletCoord, models.Coordinates{onServerData[i].BulletX, onServerData[i].BulletY, onServerData[i].IdUser})
			tankCoord = append(tankCoord, models.Coordinates{onServerData[i].TankX, onServerData[i].TankY, onServerData[i].IdUser})
			//
		}

		for b := range bulletCoord {
			for t := range tankCoord {
				if CheckHit(bulletCoord[b].X, bulletCoord[b].Y, tankCoord[t].X, tankCoord[t].Y) {
					for i := range onServerData {
						if onServerData[i].IdUser == tankCoord[t].Id {
							onServerData[i].IsAlive = false
						}
					}
				}

			}
		}

		if !slices.Contains(ids, message.IdUser) {
			onServerData = append(onServerData, message)
		}

		fmt.Printf("\n----%+v - \nonServerData----\n", onServerData)

		response := fmt.Sprintf("%+v", onServerData)
		go conn.WriteMessage(websocket.TextMessage, []byte(response))
	}
}
