<<<<<<< HEAD
package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Quantity int       `json:"quantity"`
	State    string    `json:"state"`
}

type Server struct {
	*mux.Router

	warehouseItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		warehouseItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/warehouseItems", s.createWarehouseItem()).Methods("POST")
	s.HandleFunc("/warehouseItems", s.listWarehouseItem()).Methods("GET")
	s.HandleFunc("/warehouseItems/{id}", s.removeWarehouseItem()).Methods("DELETE")
}

func (s *Server) createWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		s.warehouseItems = append(s.warehouseItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.warehouseItems); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) removeWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.warehouseItems {
			if item.ID == id {
				s.warehouseItems = append(s.warehouseItems[:i], s.warehouseItems[i+1:]...)
				break
			}
		}
	}
}
=======
package api

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Item struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Quantity int       `json:"quantity"`
	State    string    `json:"state"`
}

type Server struct {
	*mux.Router

	warehouseItems []Item
}

func NewServer() *Server {
	s := &Server{
		Router:         mux.NewRouter(),
		warehouseItems: []Item{},
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/warehouseItems", s.createWarehouseItem()).Methods("POST")
	s.HandleFunc("/warehouseItems", s.listWarehouseItem()).Methods("GET")
	s.HandleFunc("/warehouseItems/{id}", s.removeWarehouseItem()).Methods("DELETE")
}

func (s *Server) createWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i Item
		if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		i.ID = uuid.New()
		s.warehouseItems = append(s.warehouseItems, i)

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(i); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) listWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(s.warehouseItems); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}

func (s *Server) removeWarehouseItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		for i, item := range s.warehouseItems {
			if item.ID == id {
				s.warehouseItems = append(s.warehouseItems[:i], s.warehouseItems[i+1:]...)
				break
			}
		}
	}
}
>>>>>>> 5acb7b537e8f807f5d8d844c985f5e342abea78f
