<<<<<<< HEAD
package main

import (
	"net/http"

	api "github.com/RogelioMtz/Go-WarehouseServer"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
=======
package main

import (
	"net/http"

	api "github.com/RogelioMtz/Go-WarehouseServer"
)

func main() {
	srv := api.NewServer()
	http.ListenAndServe(":8080", srv)
}
>>>>>>> 5acb7b537e8f807f5d8d844c985f5e342abea78f
