package restclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
)

// FindRegisteredAddress calls the Company House API to get the Registered Office address with the input company number
func FindRegisteredAddress(companyNumber string) (roawa.RoaAddress, error) {

	log.Println("Finding Registered Addess for Company ", companyNumber, " from CH API")
	req, err := http.NewRequest("GET", "https://api.companieshouse.gov.uk/company/"+companyNumber+"/registered-office-address", nil)
	if err != nil {
		fmt.Printf("The HTTP new request failed with error %s\n", err)
		return roawa.RoaAddress{}, err
	}

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	apiKey := cfg.APIKey

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return roawa.RoaAddress{}, err
	}

	defer response.Body.Close()
	var jsonAddress roawa.RoaAddress
	err = json.NewDecoder(response.Body).Decode(&jsonAddress)
	log.Println("CH JSON Data", jsonAddress, " error value ", err)

	return jsonAddress, err
}
