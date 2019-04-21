package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DBManager struct {
	*sql.DB
}

func NewDBManager() *DBManager {
	var dm = &DBManager{
		DB: db,
	}

	return dm
}

func (dm *DBManager) StorePop(getpop *Population) error {
	res, err := dm.Query("INSERT INTO test.population (code,population) VALUES(?,?)", getpop.LocationCode, getpop.Population)
	res.Close()
	return err
}

func (dm *DBManager) StoreDis(getdis *Districts) error {
	res, err := dm.Query("INSERT INTO test.districts (en_name,code) VALUES(?,?)", getdis.Names.EnglishName, getdis.Code)
	res.Close()
	return err
}

func (dm *DBManager) StoreTown(gettown *Townships) error {
	res, err := dm.Query("INSERT INTO test.townships (tcode,dcode,en_name) VALUES(?,?,?)", gettown.Code, gettown.DistrictCode, gettown.Names.EnglishName)
	res.Close()
	return err
}

func (dm *DBManager) GetDistrictCode(name string) string {
	var code string
	res, _ := dm.Query("SELECT code from test.districts WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&code)
	}
	res.Close()
	return code
}

func (dm *DBManager) GetTownshipCodes(dtcode string) []string {
	var tcode string
	var tcodes []string
	res, _ := dm.Query("SELECT tcode from test.townships WHERE dcode=?", dtcode)
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
	res, _ := dm.Query("SELECT en_name from test.townships WHERE dcode=?", dtcode)
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
	res, _ := dm.Query("SELECT tcode from test.townships WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&tcode)
	}
	res.Close()
	return tcode
}

func (dm *DBManager) GetDistrictCodeByTownName(name string) string {
	var dcode string
	res, _ := dm.Query("SELECT dcode from test.townships WHERE en_name=?", name)
	for res.Next() {
		res.Scan(&dcode)
	}
	res.Close()
	return dcode
}
