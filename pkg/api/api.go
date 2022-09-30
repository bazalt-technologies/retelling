package api

import (
	"github.com/gorilla/mux"
	"log"
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
	api.r.HandleFunc("/api/v1/reviews", api.reviews).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/users", api.users).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
	api.r.HandleFunc("/api/v1/genres", api.genres).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/types", api.types).Methods(http.MethodGet)
	api.r.HandleFunc("/api/v1/authUser", api.authUser).Methods(http.MethodPost) // Почему http.MethodPost ?
	api.r.HandleFunc("/api/v1/content", api.content).Methods(http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodDelete)
}

func (api *API) ListenAndServe(addr string) {
	log.Print("Listen on tcp://" + addr)
	http.ListenAndServe(addr, api.r)
}
