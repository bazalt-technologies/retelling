package api

import (
	"encoding/json"
	"log"
	"net/http"
	"retelling/pkg/models"
	"strconv"
)

func (api *API) content(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		var req models.Request
		json.NewDecoder(r.Body).Decode(&req)
		data, err := api.db.GetContent(req)
		log.Println(data)
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
		var data models.Content
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = api.db.PatchContent(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	case http.MethodDelete:
		var req models.Request
		err := json.NewDecoder(r.Body).Decode(&req)
		err = api.db.DeleteContent(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
