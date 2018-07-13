package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	// Declare a router
	r := mux.NewRouter()
	// Declare static file directory
	staticFileDirectory := http.Dir("./static/")
	// Create static file server for our static files, i.e., .html, .css, etc
	staticFileServer := http.FileServer(staticFileDirectory)
	// Create file handler. Although the static files are placed inside `./static/` folder
	// in our local directory, it is served at the root (i.e., http://localhost:8080/)
	// when browsed in a browser. Hence, we need `http.StripPrefix` function to change the
	// file serve path.
	staticFileHandler := http.StripPrefix("/", staticFileServer)
	// Add static file handler to our router
	r.Handle("/", staticFileHandler).Methods("GET")
	// Add handler for `get` and `post` people functions
	r.HandleFunc("/person", getPersonHandler).Methods("GET")
	r.HandleFunc("/person", createPersonHandler).Methods("POST")

	return r
}

func main() {
	// Create router
	r := newRouter()

	// Listen to the port. Go server's default port is 8080.
	http.ListenAndServe(":8080", r)
}
