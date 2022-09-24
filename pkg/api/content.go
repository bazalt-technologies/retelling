package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) content(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var data models.Review
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		id, err := api.db.NewReview(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))
	}
}
