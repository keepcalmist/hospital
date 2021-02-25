package interfases

import "github.com/gorilla/mux"

type Connector interface {
	Router() *mux.Router
	Path() string
}
