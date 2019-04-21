package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

var (
	appname = "electroscope"
	config  Config
)

type Config struct {
	MysqlURI string
	Port     string
}

func main() {

	err := envconfig.Process(appname, &config)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	db, err := sql.Open("mysql", config.MysqlURI)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	InitStore(&dbStore{db: db})

	getdistrict()
	getpopulation()
	gettown()

	handler := http.NewServeMux()
	handler.Handle("/districts/", DistrictHandler{})
	handler.Handle("/townships/", TownshipHandler{})
	log.Fatal(http.ListenAndServe(config.Port, handler))

}
