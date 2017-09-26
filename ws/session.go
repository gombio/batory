package ws

import (
	"log"

	uuid "github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Message struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

// Session represents individial sessions connecting to WebSocket // a.k.a. SESSION
type Session struct {
	id     string
	socket *websocket.Conn
}

//Close socket connection, remove handlers, etc. from channels
func (c *Session) Close() {
	//TODO: do something
}

func (session *Session) Read() {
	var message Message
	for {
		if err := session.socket.ReadJSON(&message); nil != err {
			log.Print(session.id, ": ", err.Error())
			break
		}
		log.Print(session.id, ": ", message)
		// if handler, found := session.findHandler(message.Name); found {
		// 	handler(session, message.Data)
		// }
	}
	log.Print("Disconnecting from Reader...")
	session.socket.Close()
}

func (session *Session) Write() {
	// for msg := range session.send {
	// 	if err := session.socket.WriteJSON(msg); nil != err {
	// 		break
	// 	}
	// }
	// log.Print("Disconnecting from Writer...")
	// session.socket.Close()
}

func NewSession(socket *websocket.Conn) *Session {
	return &Session{
		id:     uuid.New().String(),
		socket: socket,
	}
}
