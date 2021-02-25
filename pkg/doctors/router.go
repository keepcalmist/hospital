package doctors

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Router struct {
	repo   Repository
	router *mux.Router
	path   string
}

func NewService(repo Repository) *Router {
	return &Router{
		repo:   repo,
		router: nil,
		path:   "/doctors",
	}
}

func (r *Router) SetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", getAllDoctors)
}
