package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) content(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	case http.MethodPost:
		var data models.Content
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		id, err := api.db.NewContent(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))
	case http.MethodPatch:
	case http.MethodDelete:
		var data models.Content
		err := json.NewDecoder(r.Body).Decode(&data)
		id, err := api.db.DeleteContent(data.ContentID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strconv.Itoa(id)))
	}
}
