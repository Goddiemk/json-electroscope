package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DistrictNames struct {
	EnglishName string `json:"en"`
}

type Districts struct {
	Code  string        `json:"code"`
	Names DistrictNames `json:"name"`
}

type TownshipNames struct {
	EnglishName string `json:"en"`
}

type Townships struct {
	Code         string        `json:"code"`
	DistrictCode string        `json:"dt_code"`
	Names        TownshipNames `json:"name"`
}

type Population struct {
	LocationCode string `json:"location_code"`
	Population   string `json:"population"`
}

var districtcodeByNames = make(map[string]string, 0)
var populationByCodes = make(map[string]int, 0)
var townshipCodeByNames = make(map[string]string, 0)
var districtCodeByTowncode = make(map[string]string, 0)

func getpopulation() {
	getpop := Population{}
	response, _ := http.Get(
		"https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/population.json",
	)
	populationData, _ := ioutil.ReadAll(response.Body)

	var populations []Population
	json.Unmarshal(populationData, &populations)
	for _, population := range populations {
		populationByCodes[population.LocationCode], _ = strconv.Atoi(population.Population)
	}
	for k, v := range populationByCodes {
		getpop.LocationCode = k
		getpop.Population = strconv.Itoa(v)
		err := dm.StorePop(&getpop)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getdistrict() {
	getdis := Districts{}
	response, _ := http.Get(
		"https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/districts.json",
	)
	districtsData, _ := ioutil.ReadAll(response.Body)

	var districts []Districts
	json.Unmarshal(districtsData, &districts)
	for _, district := range districts {
		districtcodeByNames[district.Names.EnglishName] = district.Code
	}
	for k, v := range districtcodeByNames {
		getdis.Names.EnglishName = k
		getdis.Code = v
		err := dm.StoreDis(&getdis)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func gettown() {
	gettown := Townships{}
	response, _ := http.Get("https://raw.githubusercontent.com/Electroscope/electroscope-api/master/mongo/townships.json")
	townshipsData, _ := ioutil.ReadAll(response.Body)

	var townships []Townships
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
		err := dm.StoreTown(&gettown)
		if err != nil {
			fmt.Println(err)
		}
	}

}
