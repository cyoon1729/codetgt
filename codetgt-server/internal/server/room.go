package server

import "fmt"

type Ticket struct {
	roomId   string
	userId   string
	username string
	conn     *Connection
}

type Room struct {
	roomId    string
	doc       *Document
	userConns map[string]*Connection
	userNames map[string]string

	startUserSession chan Ticket
	unregisterUser   chan string
	broadcast        chan Message
}

func CreateEmptyRoom(roomId string) *Room {
	newDoc := CreateEmptyDoc()
	newRoom := &Room{
		roomId:           roomId,
		doc:              newDoc,
		userConns:        make(map[string]*Connection),
		userNames:        make(map[string]string),
		startUserSession: make(chan Ticket),
		unregisterUser:   make(chan string),
		broadcast:        make(chan Message),
	}

	return newRoom
}

func (r *Room) register(uuid string, name string, conn *Connection) {
	r.userNames[uuid] = name
	r.userConns[uuid] = conn
	return
}

func (r *Room) unregister(uuid string) {
	close(r.userConns[uuid].send)
	delete(r.userConns, uuid)
	delete(r.userNames, uuid)
}

func (r *Room) start(uuid string) {
	fmt.Println(uuid)
	conn := r.userConns[uuid]
	go conn.writePump()
	go conn.readPump(r, uuid)
}

func (r *Room) broadcastMsg(msg Message) {
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

func (r *Room) Run() {
	for {
		select {
		case t := <-r.startUserSession:
			r.register(t.userId, t.username, t.conn)
			r.start(t.userId)
		case uuid := <-r.unregisterUser:
			r.unregister(uuid)
		case msg := <-r.broadcast:
			r.broadcastMsg(msg)
		}
	}
}
