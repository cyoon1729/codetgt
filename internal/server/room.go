package server

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

func RegisterUser(room *Room, uuid string, name string, conn *Connection) {
	room.userNames[uuid] = name
	room.userConns[uuid] = conn
	return
}

func UnregisterUser(room *Room, uuid string) {
	close(room.userConns[uuid].send)
	delete(room.userConns, uuid)
	delete(room.userNames, uuid)
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
