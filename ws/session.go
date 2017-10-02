package ws

import (
	uuid "github.com/google/uuid"
)

// Session represents individial sessions connecting to WebSocket // a.k.a. SESSION
type Session struct {
	id string
}

//Close socket connection, remove handlers, etc. from channels
func (c *Session) Close() {
	//TODO: do something
}

func NewSession() *Session {
	return &Session{
		id: uuid.New().String(),
	}
}
