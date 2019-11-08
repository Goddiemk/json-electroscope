package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"minkhantkoko/json-electroscope/lib"
)

var environ = lib.GetEnvironment()
var client = lib.GetDatabase()

func main() {
	http.HandleFunc("/districts/", districtHandler)
	http.HandleFunc("/townships/", townshipHandler)
	log.Fatal(http.ListenAndServe(environ.Port, nil))
}

func districtHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[11:]

	districtMap := make(map[string]interface{})
	districtMap["district_code"] = client.GetDistrictCode(name)
	districtMap["town_codes"] = client.GetTownshipCodes(client.GetDistrictCode(name))
	districtMap["townships"] = client.GetTownshipNames(client.GetDistrictCode(name))
	districtMap["population"] = client.GetPopulation(client.GetDistrictCode(name))
	districtJSON, err := json.Marshal(districtMap)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", districtJSON)
}

func townshipHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[11:]

	townshipMap := make(map[string]interface{})
	townshipMap["town_code"] = client.GetTownCodeByTownName(name)
	townshipMap["district_code"] = client.GetDistrictCodeByTownName(name)
	townshipMap["population"] = client.GetPopulation(client.GetTownCodeByTownName(name))
	townshipJSON, err := json.Marshal(townshipMap)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Fprintf(w, "%s", townshipJSON)
}
