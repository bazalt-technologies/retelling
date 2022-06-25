package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
)

func (api *API) newUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	_, err = api.db.NewUser(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (api *API) authUser(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	err := json.NewDecoder(r.Body).Decode(&data)
}
