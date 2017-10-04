package main

import (
	"log"
	"net/http"
	"os"

	r "github.com/GoRethink/gorethink"
	"github.com/gombio/batory/db"
	"github.com/gombio/batory/handlers"
	"github.com/gombio/batory/ws"
)

func validateFrontend() {
	if _, err := os.Stat("frontend/build/index.html"); os.IsNotExist(err) {
		log.Fatal("ERROR: Frontend root file does not exist. Make sure to run 'yarn build' first. ", err)
	}
}

func main() {
	validateFrontend()

	log.Print("Connecting to RethinkgDB...")
	session, err := r.Connect(r.ConnectOpts{
		Address:  "localhost:28015", //TODO: parameters
		Database: "batory",          //TODO: parameters
	})
	if nil != err {
		log.Fatal("Could not connect to RethinkDB", err.Error())
	}

	//TODO: verify schema
	//TODO: load schema if not already present
	//TODO: introduce "clear-db" parameter to clear schema
	rethink := db.NewRethinkdb("batory", session)
	rethink.Create(true)

	log.Print("Starting server...")

	http.Handle("/", http.FileServer(http.Dir("frontend/build/"))) //TODO: Add option to not to serve Frontend by Go

	log.Print("Setting up WebSocket server...")
	server := ws.NewServer(rethink)
	//Set up handlers
	server.Handle("schedules.list", handlers.SchedulesList)

	log.Print("Starting WebSocket server...")
	http.Handle("/ws", server)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "frontend/build/index.html")
	// })
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(hub, w, r)
	// })
	err = http.ListenAndServe("localhost:8081", nil) //TODO: Parameterize host and port
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
