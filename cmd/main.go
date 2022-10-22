package main

import (
	"github.com/gorilla/mux"
	"log"
	"retelling/pkg/api"
	"retelling/pkg/storage/pgsql"
)

const CONN = "host=127.0.0.1 port=5432 user=test password=test dbname=retelling"
const ADDR = "localhost:8081"

func main() {
	s, err := pgsql.NewStorage(CONN)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := api.New(mux.NewRouter(), s)
	a.Handle()
	a.ListenAndServe(ADDR)
}
