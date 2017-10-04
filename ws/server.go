package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gombio/batory/db"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Server handles WebSocket event routing to particular handlers
type Server struct {
	rules   map[string]Handler
	rethink *db.Rethinkdb
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil) //Upgrade normal connection to WebSocket
	if nil != err {                            //if it fails - send HTTP 500
		w.WriteHeader(http.StatusInternalServerError) //500
		fmt.Fprint(w, err.Error())
		return
	}
	session := NewSession(s.rethink) //start new client session
	log.Print("Connected: ", session.id)
	defer session.Close()       //defer session termination
	go s.Write(socket, session) //start write channel handler
	s.Read(socket, session)     //start read channel handler
}

func (s Server) Read(socket *websocket.Conn, session *Session) {
	var message Message
	for {
		if err := socket.ReadJSON(&message); nil != err {
			log.Print(session.id, ": ", err.Error())
			break
		}
		log.Print(session.id, ": ", message)
		if handler, found := s.rules[message.Type]; found {
			handler(session, message.Data)
		} else {
			log.Print("UNKNOWN MESSAGE TYPE: ", message.Type)
		}
	}
	log.Print("Disconnecting from Reader...")
	socket.Close()
}

func (s Server) Write(socket *websocket.Conn, session *Session) {
	for msg := range session.Send {
		if err := socket.WriteJSON(msg); nil != err {
			break
		}
	}
	log.Print("Disconnecting from Writer...")
	socket.Close()
}

func (s *Server) Handle(name string, handler Handler) {
	s.rules[name] = handler
}

func NewServer(rethink *db.Rethinkdb) *Server {
	return &Server{
		rules:   make(map[string]Handler),
		rethink: rethink,
	}
}
