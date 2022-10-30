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
		// Два варианта работы get дальше, поэтому decode для двух полей
		ObjectID, err := strconv.Atoi(r.URL.Query().Get("ObjectID"))
		UserID, err := strconv.Atoi(r.URL.Query().Get("UserID"))
		if ObjectID == 0 || err != nil {
			json.NewDecoder(r.Body).Decode(&req)
		}
		if UserID == 0 || err != nil {
			json.NewDecoder(r.Body).Decode(&req)
		}
		req.ObjectID = ObjectID
		req.UserID = UserID
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
