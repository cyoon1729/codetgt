package server

import (
	"fmt"
	"log"
	"net/http"

	"codetgt/internal/session"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Lobby struct {
	Rooms      map[string]map[*session.Room]bool
	Register   chan session.Ticket
	Unregister chan session.Ticket
}

type ConnInfo struct {
	RoomId string
	Uuid   string
	Name   string
}

func (lobby *Lobby) Run() {
	for {
		select {
		case t := <-lobby.Register:
			fmt.Println(t.Username)
		}
	}
}

func EnterRoom(lobby *Lobby, w http.ResponseWriter, r *http.Request, usr ConnInfo) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	conn := &session.Connection{Ws: ws, Send: make(chan []byte, 256)}
	ticket := session.Ticket{
		RoomId:   usr.RoomId,
		UserId:   usr.Uuid,
		Username: usr.Name,
		Conn:     conn,
	}
	lobby.Register <- ticket
	return
}

func doshit(c *session.Connection) {
	doshit(c)
}
