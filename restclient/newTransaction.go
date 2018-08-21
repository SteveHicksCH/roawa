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

// CreateTransaction calls CH to create a new Transaction for future API calls
func CreateTransaction(transactionRequest roawa.TransactionRequest) (roawa.TransactionResponse, error) {
	log.Println("Getting a new transaction", transactionRequest)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	apiKey := cfg.LocalAPIKey
	transactionURL := cfg.LocalTxnAddress

	jsonRequestValue, err := json.Marshal(transactionRequest)
	if err != nil {
		fmt.Printf("Fail to marshall request  %s\n", err)
		return roawa.TransactionResponse{}, err
	}

	req, err := http.NewRequest("POST", transactionURL, bytes.NewBuffer(jsonRequestValue))
	if err != nil {
		log.Printf("The HTTP new transaction request failed with error %s\n", err)
		return roawa.TransactionResponse{}, err
	}

	req.SetBasicAuth(apiKey, "")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
		return roawa.TransactionResponse{}, err
	}
	if response.StatusCode != http.StatusCreated {
		msg := fmt.Sprintf("Wrong Http Status Code (must be 201), reason:  %s\n", response.Status)
		log.Printf(msg)
		return roawa.TransactionResponse{}, errors.New(msg)
	}

	defer response.Body.Close()
	var jsonTransactionResponse roawa.TransactionResponse
	err = json.NewDecoder(response.Body).Decode(&jsonTransactionResponse)
	log.Println("Txn Response JSON", jsonTransactionResponse, " error value ", err)

	return jsonTransactionResponse, err
}
