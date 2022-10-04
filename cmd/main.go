package main

import (
	"github.com/gorilla/mux"
	"log"
	"retelling/pkg/api"
	"retelling/pkg/storage/pgsql"
)

func main() {
	s, err := pgsql.NewStorage(CONN)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := api.New(mux.NewRouter(), s)
	a.Handle()
	a.ListenAndServe(ADDR)
}
