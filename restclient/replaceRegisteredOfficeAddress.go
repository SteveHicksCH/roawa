package restclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
)

// ReplaceROA replaces the current registered office address for the company
func ReplaceROA(address roawa.RoaAddress, transactionID string) error {
	log.Println("Updating the registered office address with ", address, " with transaction ", transactionID)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	apiKey := cfg.LocalAPIKey
	replaceROAURL := cfg.LocalROAAddressPrefix + transactionID + "/registered-office-address"

	jsonRequestValue, err := json.Marshal(address)
	if err != nil {
		fmt.Printf("Fail to marshall request  %s\n", err)
		return err
	}

	req, err := http.NewRequest("PUT", replaceROAURL, bytes.NewBuffer(jsonRequestValue))
	if err != nil {
		log.Printf("The HTTP replace ROA request failed with error %s\n", err)
		return err
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return err
	}
	if !(response.StatusCode == http.StatusNoContent || response.StatusCode == http.StatusOK) {
		msg := fmt.Sprintf("Wrong Http Status Code (must be 200 or 204), reason:  %s\n", response.Status)
		log.Printf(msg)
		return errors.New(msg)
	}

	return nil
}
