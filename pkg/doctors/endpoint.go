package doctors

import "net/http"

func (r *Router) getAllDoctors(w http.ResponseWriter, req *http.Request) {
	doctors, err := r.repo.SelectAll(req.Context())
	if err != nil {
		return
	}

}
