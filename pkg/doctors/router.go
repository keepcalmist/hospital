package doctors

import (
	"github.com/gorilla/mux"
	"github.com/keepcalmist/hospital/pkg/interfases"
	"net/http"
)

type DoctorService struct {
	repo   Repository
	router *mux.Router
	path   string
}

func NewService(repo Repository) interfases.Connector {

	dcts := &DoctorService{
		repo:   repo,
		router: nil,
		path:   "/doctors",
	}
	dcts.SetRouter()
	return dcts
}

func (d *DoctorService) Router() *mux.Router {
	return d.router
}

func (d *DoctorService) Path() string {
	return d.path
}

func (r *DoctorService) SetRouter() {
	router := mux.NewRouter()
	router.HandleFunc("/doctors", r.getAllDoctors).Methods(http.MethodGet)
	router.Handle("/actor", addLoshok())

	r.router = router
}
