package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
)

func (api *API) likes(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var req models.Request
		req.ObjectID = paramInt(r, "ObjectID")
		req.UserID = paramInt(r, "UserID")
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
	}
}
