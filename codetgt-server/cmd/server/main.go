package main

import (
	"codetgt-server/internal/server"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!")

	lobby := &server.Lobby{
		Rooms:      make(map[string]*server.Room),
		CreateRoom: make(chan string),
		Register:   make(chan server.Ticket),
		Unregister: make(chan server.Ticket),
	}

	go lobby.Run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
		roomId := c.Param("roomId")
		server.SpawnRoom(lobby, roomId)
	})

	router.GET("/ws/:roomId/:uuid/:name", func(c *gin.Context) {
		usr := server.ConnInfo{
			RoomId: c.Param("roomId"),
			Uuid:   c.Param("uuid"),
			Name:   c.Param("name"),
		}
		server.EnterRoom(lobby, c.Writer, c.Request, usr)
	})

	router.Run("0.0.0.0:8888")
}
