package session

type Room struct {
	roomId    string
	document  Document
	userConns map[string]map[*Connection]bool
	userNames map[string]string
}

type Ticket struct {
	RoomId   string
	UserId   string
	Username string
	Conn     *Connection
}

func (room *Room) Run() {
	return
}
