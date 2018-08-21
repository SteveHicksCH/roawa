package restclient

import (
	"fmt"
	"testing"

	"github.com/shicks/roawa"
)

func TestTransaction(t *testing.T) {
	fmt.Println("Running create Transaction test")
	transactionRequest := roawa.TransactionRequest{
		CompanyNumber: "00006400",
		Reference:     "",
		Description:   "description",
	}
	transactionResponse, err := CreateTransaction(transactionRequest)

	fmt.Println(transactionResponse, err)

}
