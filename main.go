package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	InitStore(&dbStore{db: db})

	//getdistrict()
	//getpopulation()
	//gettown()

	handler := http.NewServeMux()
	handler.Handle("/districts/", DistrictHandler{})
	handler.Handle("/townships/", TownshipHandler{})
	log.Fatal(http.ListenAndServe(":8080", handler))

}
