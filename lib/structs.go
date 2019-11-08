package lib

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
