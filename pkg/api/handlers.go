package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) reviews(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req := models.Request{UserID: id}
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
}

func (api *API) newUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = api.db.NewUser(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (api *API) newReview(w http.ResponseWriter, r *http.Request) {
	var data models.Review
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = api.db.NewReview(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
