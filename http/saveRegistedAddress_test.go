package http

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	fmt.Println("Running Save Test")
	w := httptest.NewRecorder()

	data := url.Values{}
	data.Set("premises", "16")
	data.Add("postcode", "CF64 2SQ")
	data.Add("addressLine1", "Coleridge Avenue")
	data.Add("addressLine2", "The Gardens")
	data.Add("county", "Vale of Glamorgan")
	data.Add("country", "uk")
	data.Add("town", "Penarth")

	reader := strings.NewReader(data.Encode())
	// create a fake request to be passed into the handler
	r, err := http.NewRequest("POST", urlPrefix+"/save", reader)
	if err != nil {
		t.Fatalf("error constructing test HTTP request [%s]", err)
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	saveRegistedAddress(w, r)

	scanner := bufio.NewScanner(w.Body)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "New Address") {
			newAddressHTMLLine := scanner.Text()
			expectedAddress := "New Address  16 Coleridge Avenue The Gardens Penarth Vale of Glamorgan CF64 2SQ United Kingdom"
			if !strings.Contains(newAddressHTMLLine, expectedAddress) {
				t.Errorf("Incorrect new address, expected = '%s' , within line '%s'", expectedAddress, newAddressHTMLLine)
			}
		}
	}

}
