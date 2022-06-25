package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) newUser(w http.ResponseWriter, r *http.Request) {
	var data models.User
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	id, err := api.db.NewUser(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write([]byte(strconv.Itoa(id)))
}

func (api *API) authUser(w http.ResponseWriter, r *http.Request) {
	var req models.Request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	id, err := api.db.AuthUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
	w.Write([]byte(strconv.Itoa(id)))
}
