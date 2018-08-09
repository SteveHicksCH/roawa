package roawa

var countries *map[string]string

// GetCountries Populate a list of countries
func GetCountries() *map[string]string {

	if countries == nil {
		countries = &map[string]string{
			"england":  "England",
			"ni":       "Northern Ireland",
			"scotland": "Scotland",
			"uk":       "United Kingdom",
			"wales":    "Wales",
			"ns":       "Not Specified",
		}
	}
	return countries
}
