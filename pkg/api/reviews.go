package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
)

func (api *API) reviews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		var req models.Request
		req.UserID = paramInt(r, "UserID")

		data, err := api.db.GetReviews(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)

	case http.MethodPost:
		var data models.Review
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		_, err = api.db.NewReview(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
	}
}
