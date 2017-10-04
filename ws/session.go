package ws

import (
	"github.com/gombio/batory/db"
	uuid "github.com/google/uuid"
)

// Session represents individial sessions connecting to WebSocket // a.k.a. SESSION
type Session struct {
	id      string
	Send    chan Message
	Rethink *db.Rethinkdb
}

//Close socket connection, remove handlers, etc. from channels
func (c *Session) Close() {
	//TODO: do something
}

func NewSession(Rethink *db.Rethinkdb) *Session {
	return &Session{
		id:      uuid.New().String(),
		Send:    make(chan Message),
		Rethink: Rethink,
	}
}
