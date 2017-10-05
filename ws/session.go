package ws

import (
	"github.com/gombio/batory/db"
	uuid "github.com/google/uuid"
)

// Session represents individial sessions connecting to WebSocket // a.k.a. SESSION
type Session struct {
	ID              string
	Send            chan Message
	Rethink         *db.Rethinkdb
	User            db.User
	IsAuthenticated bool
}

func (s Session) Identifier() string {
	if s.IsAuthenticated {
		return s.User.ID
	}

	return s.ID
}

//Close socket connection, remove handlers, etc. from channels
func (c *Session) Close() {
	//TODO: do something
}

func NewSession(Rethink *db.Rethinkdb) *Session {
	return &Session{
		ID:              uuid.New().String(),
		Send:            make(chan Message),
		Rethink:         Rethink,
		IsAuthenticated: false,
	}
}
