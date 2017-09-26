package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gombio/batory/ws"
)

func validateFrontend() {
	if _, err := os.Stat("frontend/build/index.html"); os.IsNotExist(err) {
		log.Fatal("ERROR: Frontend root file does not exist. Make sure to run 'yarn build' first. ", err)
	}
}

func main() {
	validateFrontend()
	log.Print("Starting server...")

	http.Handle("/", http.FileServer(http.Dir("frontend/build/"))) //TODO: Add option to not to serve Frontend by Go

	log.Print("Setting up WebSocket Router...")
	router := ws.NewRouter()

	http.Handle("/ws", router)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "frontend/build/index.html")
	// })
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(hub, w, r)
	// })
	err := http.ListenAndServe("localhost:8080", nil) //TODO: Parameterize host and port
	if err != nil {
		log.Fatal("Could not start server: ", err)
	}
}
