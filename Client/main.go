// main.go
// This is the main entry point for the GoTestServer application.
// It sets up an HTTP server using the Gorilla Mux router and listens on localhost at port 8080.

package main

import (
	"net/http"

	"github.com/RogelioMtz/server"
)

func main() {
	srv := server.NewServer()
	http.ListenAndServe(":8080", srv)
}
