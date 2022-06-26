package session

type Room struct {
	RoomId    string
	Doc       *Document
	UserConns map[string]*Connection
	UserNames map[string]string
}

type Ticket struct {
	RoomId   string
	UserId   string
	Username string
	Conn     *Connection
}

func CreateEmptyRoom(roomId string) *Room {
	newDoc := CreateEmptyDoc()
	newRoom := &Room{
		RoomId:    roomId,
		Doc:       newDoc,
		UserConns: make(map[string]*Connection),
		UserNames: make(map[string]string),
	}

	return newRoom
}

func RegisterUser(room *Room, uuid string, name string, conn *Connection) {
	room.UserNames[uuid] = name
	room.UserConns[uuid] = conn
}

func UnregisterUser(room *Room, uuid string) {
	close(room.UserConns[uuid].Send)
	delete(room.UserConns, uuid)
	delete(room.UserNames, uuid)
}

func (room *Room) Run() {
	return
}
