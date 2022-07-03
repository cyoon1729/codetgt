package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Lobby struct {
	Rooms      map[string]*Room
	CreateRoom chan string
	Register   chan Ticket
	Unregister chan Ticket
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
			newRoom := CreateEmptyRoom(roomId)
			lobby.Rooms[roomId] = newRoom
			go lobby.Rooms[roomId].Run()
		case t := <-lobby.Register:
			room := lobby.Rooms[t.roomId]
			room.startUserSession <- t
		case t := <-lobby.Unregister:
			room := lobby.Rooms[t.roomId]
			room.unregisterUser <- t.userId
		}
	}
}

func SpawnRoom(lobby *Lobby, roomId string) {
	_, ok := lobby.Rooms[roomId]
	if !ok {
		lobby.CreateRoom <- roomId
	}
	return
}

func EnterRoom(lobby *Lobby, w http.ResponseWriter, r *http.Request, usr ConnInfo) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	conn := &Connection{ws: ws, send: make(chan []byte, 256)}
	conn.initSocketConn()

	ticket := Ticket{
		roomId:   usr.RoomId,
		userId:   usr.Uuid,
		username: usr.Name,
		conn:     conn,
	}

	lobby.Register <- ticket
}
