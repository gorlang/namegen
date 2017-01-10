package main

/*
	Gorilla mux is used to handle requests paths and routing.
	https://github.com/gorilla/mux
	CORS added to be able to use it from another port number for standalone testing purposes only!
	(Node.js or webpage executing on you usual webserver)
	(DO NOT USE IN PRODUCTION)
	Instead you could for example configure your http server as a reverse proxy to relay request from :80 to the
	backend service.
*/

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/names", NamesHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS()(r)))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HomeHandler!\n"))
}
