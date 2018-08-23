package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var store Store

func InitStore(s Store) {
	store = s
}

type Store interface {
	StorePop(getpop *Population) error
	StoreDis(getdis *Districts) error
	StoreTown(gettown *Townships) error
	GetDistrictCode(name string) string
	GetTownshipCodes(dtcode string) []string
	GetTownshipNames(dtcode string) []string
	GetPopulation(code string) string
	GetTownCodeByTownName(name string) string
	GetDistrictCodeByTownName(name string) string
}

type dbStore struct {
	db *sql.DB
}

func (store *dbStore) StorePop(getpop *Population) error {

	res, err := store.db.Query("INSERT INTO test.population (code,population) VALUES('" + getpop.LocationCode + "','" + getpop.Population + "')")
	res.Close()
	return err
}

func (store *dbStore) StoreDis(getdis *Districts) error {
	res, err := store.db.Query("INSERT INTO test.districts (en_name,code) VALUES('" + getdis.Names.EnglishName + "','" + getdis.Code + "')")
	res.Close()
	return err
}

func (store *dbStore) StoreTown(gettown *Townships) error {
	res, err := store.db.Query("INSERT INTO test.townships (tcode,dcode,en_name) VALUES('" + gettown.Code + "','" + gettown.DistrictCode + "','" + gettown.Names.EnglishName + "')")
	res.Close()
	return err
}

func (store *dbStore) GetDistrictCode(name string) string {
	var code string
	res, _ := store.db.Query("SELECT code from test.districts WHERE en_name='" + name + "'")
	for res.Next() {
		res.Scan(&code)
	}
	res.Close()
	return code
}

func (store *dbStore) GetTownshipCodes(dtcode string) []string {
	var tcode string
	var tcodes []string

	res, _ := store.db.Query("SELECT tcode from test.townships WHERE dcode='" + dtcode + "'")
	for res.Next() {
		res.Scan(&tcode)
		tcodes = append(tcodes, tcode)
	}
	res.Close()
	return tcodes
}

func (store *dbStore) GetTownshipNames(dtcode string) []string {
	var tname string
	var tnames []string
	res, _ := store.db.Query("SELECT en_name from test.townships WHERE dcode='" + dtcode + "'")
	for res.Next() {
		res.Scan(&tname)
		tnames = append(tnames, tname)
	}

	res.Close()
	return tnames
}

func (store *dbStore) GetPopulation(code string) string {
	var population string
	res, _ := store.db.Query("SELECT population from test.population WHERE code='" + code + "'")
	for res.Next() {
		res.Scan(&population)
	}
	res.Close()
	return population
}

func (store *dbStore) GetTownCodeByTownName(name string) string {
	var tcode string
	res, _ := store.db.Query("SELECT tcode from test.townships WHERE en_name='" + name + "'")
	for res.Next() {
		res.Scan(&tcode)
	}
	res.Close()
	return tcode
}

func (store *dbStore) GetDistrictCodeByTownName(name string) string {
	var dcode string
	res, _ := store.db.Query("SELECT dcode from test.townships WHERE en_name='" + name + "'")
	for res.Next() {
		res.Scan(&dcode)
	}
	res.Close()
	return dcode
}
