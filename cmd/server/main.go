package main

import (
	"codetgt/internal/server"
	"codetgt/internal/session"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world!")

	lobby := &server.Lobby{
		Rooms:      make(map[string]*session.Room),
		CreateRoom: make(chan string),
		Register:   make(chan session.Ticket),
		Unregister: make(chan session.Ticket),
	}

	go lobby.Run()

	router := gin.New()
	router.LoadHTMLFiles("index.html")

	router.GET("/room", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
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
