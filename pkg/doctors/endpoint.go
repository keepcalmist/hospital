package doctors

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/keepcalmist/hospital/pkg/tools/encodeReponses"
	"net/http"
)

func (r *DoctorService) getAllDoctors(w http.ResponseWriter, req *http.Request) {
	doctors, err := r.repo.SelectAll(req.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, err := encodeReponses.Json(doctors)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
	w.WriteHeader(http.StatusOK)
	return
}

func addLoshok() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/lol", lol)
	r.HandleFunc("/actor/nelol", nelol)
	return r
}

func lol(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "lol")
}

func nelol(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "nelol")
}
