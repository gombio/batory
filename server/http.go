package server

import (
	"fmt"
	"github.com/codegangsta/negroni"
	//"github.com/meatballhat/negroni-logrus"
	"net/http"
)

type httpServer struct {
	address string
	port    string
}

func NewHttpServer(add, port string) *httpServer {
	return &httpServer{address: add, port: port}
}

func (server *httpServer) Execute() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/twilio/incoming_call", callHandler)
	n := negroni.Classic()
	//n.Use(negronilogrus.NewMiddleware())
	add := fmt.Sprintf("%s:%s", server.address, server.port)
	n.UseHandler(mux)
	n.Run(add)
}
