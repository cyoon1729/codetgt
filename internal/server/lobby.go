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
	Rooms      map[string]*session.Room
	CreateRoom chan string
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
		case roomId := <-lobby.CreateRoom:
			newRoom := session.CreateEmptyRoom(roomId)
			lobby.Rooms[roomId] = newRoom
		case t := <-lobby.Register:
			fmt.Println(t.Username)
			room := lobby.Rooms[t.RoomId]
			session.RegisterUser(room, t.UserId, t.Username, t.Conn)
		case t := <-lobby.Unregister:
			fmt.Println(t.Username)
			room := lobby.Rooms[t.RoomId]
			session.UnregisterUser(room, t.UserId)
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
