package api

import (
	"encoding/json"
	"net/http"
	"retelling/pkg/models"
)

func (api *API) types(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		var req models.Request
		req.ObjectID = paramInt(r, "ObjectID") //

		data, err := api.db.GetTypes(req)
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
	}
}
