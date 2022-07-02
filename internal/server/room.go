package server

import "fmt"

type Room struct {
	roomId    string
	doc       *Document
	userConns map[string]*Connection
	userNames map[string]string
}

type Ticket struct {
	roomId   string
	userId   string
	username string
	conn     *Connection
}

func CreateEmptyRoom(roomId string) *Room {
	newDoc := CreateEmptyDoc()
	newRoom := &Room{
		roomId:    roomId,
		doc:       newDoc,
		userConns: make(map[string]*Connection),
		userNames: make(map[string]string),
	}

	return newRoom
}

func (r *Room) registerUser(uuid string, name string, conn *Connection) {
	r.userNames[uuid] = name
	r.userConns[uuid] = conn
	return
}

func (r *Room) unregisterUser(uuid string) {
	close(r.userConns[uuid].send)
	delete(r.userConns, uuid)
	delete(r.userNames, uuid)
}

func (r *Room) startUserSession(uuid string) {
	fmt.Println(uuid)
	conn := r.userConns[uuid]
	go conn.writePump()
	go conn.readPump()
}

func (r *Room) Broadcast(msg Message) {
	for usr := range r.userConns {
		c := r.userConns[usr]
		select {
		case c.send <- msg.data:
		default:
			close(c.send)
			delete(r.userConns, usr)
		}
	}
}

func (room *Room) Run() {
	return
}
