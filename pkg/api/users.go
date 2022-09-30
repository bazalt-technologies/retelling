package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		var req models.Request
		req.ObjectID = paramInt(r, "ObjectID") //

		data, err := api.db.GetUsers(req)
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
		var data models.User
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		id, err := api.db.NewUser(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))

	case http.MethodPatch:
		var data models.User
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		id, err := api.db.UpdateUser(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))

	case http.MethodDelete:
		var data models.User
		err := json.NewDecoder(r.Body).Decode(&data)
		id, err := api.db.DeleteUser(data.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))
	}
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
