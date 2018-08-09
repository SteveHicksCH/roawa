package restclient

import (
	"fmt"
	"testing"

	"github.com/shicks/roawa"
)

func TestFindRegisteredAddress(t *testing.T) {
	fmt.Println("Running find registered address test")
	var address roawa.RoaAddress
	var err error
	address, err = FindRegisteredAddress("00000006")

	fmt.Println(address, err)

}
