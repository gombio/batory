package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func testHandler(x interface{}) {
	log.Print("testHandler")
	log.Print(x)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Router handles WebSocket event routing to particular handlers
type Router struct {
	rules map[string]func(interface{})
}

func (e Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	socket, err := upgrader.Upgrade(w, r, nil)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError) //500
		fmt.Fprint(w, err.Error())
		return
	}
	session := NewSession(socket)
	log.Print("Connected: ", session.id)
	defer session.Close()
	go session.Write()
	session.Read()
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]func(interface{})),
	}
}
