package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type DistrictHandler struct{}

func (districthandler DistrictHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Path[11:]

	districtMap := make(map[string]interface{})
	districtMap["district_code"] = store.GetDistrictCode(name)
	districtMap["town_codes"] = store.GetTownshipCodes(store.GetDistrictCode(name))
	districtMap["townships"] = store.GetTownshipNames(store.GetDistrictCode(name))
	districtMap["population"] = store.GetPopulation(store.GetDistrictCode(name))
	districtJSON, err := json.Marshal(districtMap)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", districtJSON)

}

type TownshipHandler struct{}

func (townshiphandler TownshipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[11:]

	townshipMap := make(map[string]interface{})
	townshipMap["town_code"] = store.GetTownCodeByTownName(name)
	townshipMap["district_code"] = store.GetDistrictCodeByTownName(name)
	townshipMap["population"] = store.GetPopulation(store.GetTownCodeByTownName(name))
	townshipJSON, err := json.Marshal(townshipMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", townshipJSON)
}
