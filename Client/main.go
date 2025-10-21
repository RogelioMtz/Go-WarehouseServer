package main

import (
	"net/http"

	api "github.com/RogelioMtz/Go-WarehouseServer"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
