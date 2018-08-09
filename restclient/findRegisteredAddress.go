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

	req, err := http.NewRequest("GET", "https://api.companieshouse.gov.uk/company/"+companyNumber+"/registered-office-address", nil)
	if err != nil {
		fmt.Printf("The HTTP new request failed with error %s\n", err)
		return roawa.RoaAddress{}, err
	}
	// read in username

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
	var jsonAddress roawa.RoaAddress

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return roawa.RoaAddress{}, err
	} else {
		defer response.Body.Close()
		err = json.NewDecoder(response.Body).Decode(&jsonAddress)
		log.Println(jsonAddress)
	}

	return jsonAddress, nil
}
