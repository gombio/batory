package handlers

import (
	"log"

	"github.com/gombio/batory/ws"
	// r "github.com/GoRethink/gorethink"
)

func EventsList(s *ws.Session, data interface{}) {
	log.Print(s)
	log.Print(data)
}
