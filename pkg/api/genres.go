package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
)

func (api *API) genres(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		var req models.Request
		json.NewDecoder(r.Body).Decode(&req)
		data, err := api.db.GetGenres(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
