package Login

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"net/http"
)

func MakeLoginHandler() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/login", login)
	return r
}

func login(w http.ResponseWriter, r *http.Request) {
	decoder := schema.NewDecoder()
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var login Login

	err = decoder.Decode(&login, r.PostForm)

}
