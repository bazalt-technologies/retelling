package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"retelling/pkg/storage"
)

type API struct {
	r  *mux.Router
	db storage.DB
}

func New(r *mux.Router, s storage.DB) *API {
	return &API{r: r, db: s}
}

// Handle API paths.
func (api *API) Handle() {
	api.r.HandleFunc("/api/v1/getReviews/{id}", api.reviews).Methods(http.MethodGet, http.MethodOptions)
	api.r.HandleFunc("/api/v1/newUser", api.newUser).Methods(http.MethodPost, http.MethodOptions)
	api.r.HandleFunc("/api/v1/newReview", api.newReview).Methods(http.MethodPost, http.MethodOptions)
}
