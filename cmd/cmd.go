package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"minkhantkoko/json-electroscope/lib"

	"github.com/trhura/simplecli"
)

type CmdOperations struct{}

func (c CmdOperations) Create() {
	db := lib.GetDatabase()
	db.CreateTables()
}

func (c CmdOperations) Seed() {
	var districtcodeByNames = make(map[string]string)
	var populationByCodes = make(map[string]int)
	var townshipCodeByNames = make(map[string]string)
	var districtCodeByTowncode = make(map[string]string)

	var db = lib.GetDatabase()
	getpop := lib.Population{}
	response, _ := http.Get(
		"https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/population.json",
	)
	populationData, _ := ioutil.ReadAll(response.Body)

	var populations []lib.Population
	json.Unmarshal(populationData, &populations)
	for _, population := range populations {
		populationByCodes[population.LocationCode], _ = strconv.Atoi(population.Population)
	}
	for k, v := range populationByCodes {
		getpop.LocationCode = k
		getpop.Population = strconv.Itoa(v)
		err := db.StorePop(&getpop)
		if err != nil {
			fmt.Println(err)
		}
	}

	getdis := lib.Districts{}
	response, _ = http.Get(
		"https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/districts.json",
	)
	districtsData, _ := ioutil.ReadAll(response.Body)

	var districts []lib.Districts
	json.Unmarshal(districtsData, &districts)
	for _, district := range districts {
		districtcodeByNames[district.Names.EnglishName] = district.Code
	}
	for k, v := range districtcodeByNames {
		getdis.Names.EnglishName = k
		getdis.Code = v
		err := db.StoreDis(&getdis)
		if err != nil {
			fmt.Println(err)
		}
	}

	gettown := lib.Townships{}
	response, _ = http.Get("https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/townships.json")
	townshipsData, _ := ioutil.ReadAll(response.Body)

	var townships []lib.Townships
	json.Unmarshal(townshipsData, &townships)

	for _, town := range townships {
		districtCodeByTowncode[town.Code] = town.DistrictCode
	}

	for _, town := range townships {
		townshipCodeByNames[town.Names.EnglishName] = town.Code
	}

	for k, v := range districtCodeByTowncode {
		gettown.Code = k
		gettown.DistrictCode = v
		for k, v := range townshipCodeByNames {
			if v == gettown.Code {
				gettown.Names.EnglishName = k
				gettown.Code = v
			}

		}
		err := db.StoreTown(&gettown)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	simplecli.Handle(&CmdOperations{})
}
