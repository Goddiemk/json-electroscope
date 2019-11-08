package lib

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	*sql.DB
}

var dbOnce sync.Once
var dbInstance *DBManager

func GetDatabase() *DBManager {
	dbOnce.Do(func() {
		var db *sql.DB
		var err error

		if db, err = sql.Open("mysql", environ.MYSQLURI); err != nil {
			fmt.Println(err)
		}

		dbInstance = &DBManager{
			DB: db,
		}
	})
	return dbInstance
}

var populationTableSQL = `
	CREATE TABLE population (
		id	SERIAL,
		code VARCHAR(20),
		population VARCHAR(20)
	);`

var districtTableSQL = `
	CREATE TABLE district (
		id SERIAL,
		code VARCHAR(20),
		en_name VARCHAR(20)
	);`

var townshipTableSQL = `
	CREATE TABLE township(
		id SERIAL,
		tcode VARCHAR(20),
		dcode VARCHAR(20),
		en_name VARCHAR(20)
	)`

func (dm *DBManager) StorePop(getpop *Population) error {
	_, err := dm.Exec("INSERT INTO population (code,population) VALUES(?,?)", getpop.LocationCode, getpop.Population)
	return err
}

func (dm *DBManager) StoreDis(getdis *Districts) error {
	_, err := dm.Exec("INSERT INTO district (en_name,code) VALUES(?,?)", getdis.Names.EnglishName, getdis.Code)
	return err
}

func (dm *DBManager) StoreTown(gettown *Townships) error {
	_, err := dm.Exec("INSERT INTO township (tcode,dcode,en_name) VALUES(?,?,?)", gettown.Code, gettown.DistrictCode, gettown.Names.EnglishName)
	return err
}

func (dm *DBManager) GetDistrictCode(name string) string {
	var code string
	res, _ := dm.Query("SELECT code from district WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&code)
	}
	res.Close()
	return code
}

func (dm *DBManager) GetTownshipCodes(dtcode string) []string {
	var tcode string
	var tcodes []string
	res, _ := dm.Query("SELECT tcode from township WHERE dcode=?", dtcode)
	for res.Next() {
		res.Scan(&tcode)
		tcodes = append(tcodes, tcode)
	}
	res.Close()
	return tcodes
}

func (dm *DBManager) GetTownshipNames(dtcode string) []string {
	var tname string
	var tnames []string
	res, _ := dm.Query("SELECT en_name from township WHERE dcode=?", dtcode)
	for res.Next() {
		res.Scan(&tname)
		tnames = append(tnames, tname)
	}

	res.Close()
	return tnames
}

func (dm *DBManager) GetPopulation(code string) string {
	var population string
	res, _ := dm.Query("SELECT population from test.population WHERE code=?", code)
	for res.Next() {
		res.Scan(&population)
	}
	res.Close()
	return population
}

func (dm *DBManager) GetTownCodeByTownName(name string) string {
	var tcode string
	res, _ := dm.Query("SELECT tcode from test.township WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&tcode)
	}
	res.Close()
	return tcode
}

func (dm *DBManager) GetDistrictCodeByTownName(name string) string {
	var dcode string
	res, _ := dm.Query("SELECT dcode from test.township WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&dcode)
	}
	res.Close()
	return dcode
}

func (dm *DBManager) CreateTables() error {
	if _, err := dm.Exec(populationTableSQL); err != nil {
		return err
	}

	if _, err := dm.Exec(districtTableSQL); err != nil {
		return err
	}

	if _, err := dm.Exec(townshipTableSQL); err != nil {
		return err
	}
	return nil
}
