package main

import (
	"github.com/gorilla/mux"
	"log"
	"retelling/cmd/config"
	"retelling/pkg/api"
	"retelling/pkg/storage/pgsql"
)

func main() {
	s, err := pgsql.NewStorage(config.CONN)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := api.New(mux.NewRouter(), s)
	a.Handle()
	a.ListenAndServe(config.ADDR)
}
