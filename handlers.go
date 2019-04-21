package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type DistrictHandler struct{}

func (DistrictHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Path[11:]

	districtMap := make(map[string]interface{})
	districtMap["district_code"] = dm.GetDistrictCode(name)
	districtMap["town_codes"] = dm.GetTownshipCodes(dm.GetDistrictCode(name))
	districtMap["townships"] = dm.GetTownshipNames(dm.GetDistrictCode(name))
	districtMap["population"] = dm.GetPopulation(dm.GetDistrictCode(name))
	districtJSON, err := json.Marshal(districtMap)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", districtJSON)

}

type TownshipHandler struct{}

func (TownshipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[11:]

	townshipMap := make(map[string]interface{})
	townshipMap["town_code"] = dm.GetTownCodeByTownName(name)
	townshipMap["district_code"] = dm.GetDistrictCodeByTownName(name)
	townshipMap["population"] = dm.GetPopulation(dm.GetTownCodeByTownName(name))
	townshipJSON, err := json.Marshal(townshipMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", townshipJSON)
}
