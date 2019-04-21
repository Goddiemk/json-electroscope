package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var (
	db     *sql.DB
	config Config
	dm     *DBManager
	err    error
)

type Config struct {
	MysqlURI string
	Port     string
}

func init() {
	if err = godotenv.Load("system.env"); err != nil {
		fmt.Println(err.Error())
	}

	err := envconfig.Process("", &config)
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err = sql.Open("mysql", config.MysqlURI)
	if err != nil {
		fmt.Println(err)
	}

	dm = NewDBManager()
}

func main() {

	getdistrict()
	getpopulation()
	gettown()

	handler := http.NewServeMux()
	handler.Handle("/districts/", DistrictHandler{})
	handler.Handle("/townships/", TownshipHandler{})
	log.Fatal(http.ListenAndServe(config.Port, handler))

}
