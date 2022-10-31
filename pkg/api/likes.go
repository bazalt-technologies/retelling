package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) likes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var req models.Request
		id, err := strconv.Atoi(r.URL.Query().Get("ObjectID"))
		uId, err := strconv.Atoi(r.URL.Query().Get("UserID"))
		if id == 0 || err != nil || uId == 0 {
			json.NewDecoder(r.Body).Decode(&req)
		}
		req.ObjectID = id
		req.UserID = uId
		if req.UserID != 0 {
			// Лайки пользователя
			data, err := api.db.GetLikes(req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = json.NewEncoder(w).Encode(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			// Пользователи, лайкнувшие контент
			data, err := api.db.GetUsersLiked(req)
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
	case http.MethodPost:
		var data models.Request
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.NewLike(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case http.MethodDelete:
		var data models.Request
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.Unlike(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
