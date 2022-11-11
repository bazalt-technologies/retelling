package main

import (
	"github.com/gorilla/mux"
	"log"
	"retelling/pkg/api"
	"retelling/pkg/storage/pgsql"
)

const CONN = "host=database port=5432 user=test password=test dbname=retelling"
const ADDR = "server:8081"

func main() {
	s, err := pgsql.NewStorage(CONN)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := api.New(mux.NewRouter(), s)
	a.Handle()
	a.ListenAndServe(ADDR)
}
