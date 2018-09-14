package http

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
	"github.com/shicks/roawa/restclient"
)

// editRegistedAddress is a handler for when a user starts to edit the company registered address.EditRegistedAddress.EditRegistedAddress
// It populates the Edit Registered Address Screen
func editRegistedAddress(w http.ResponseWriter, r *http.Request) {

	// Get companyName from the request
	keys, ok := r.URL.Query()["companyID"]
	if !ok || len(keys[0]) < 1 {
		log.Println("No companyID passed - redirect to company ID entry page ", findURL)
		http.Redirect(w, r, findURL, http.StatusTemporaryRedirect)
		return
	}

	fmt.Println("keys", keys, ok)

	CompanyID := keys[0]
	log.Println("Company ID entered", CompanyID)

	roaAddress, err := restclient.FindRegisteredAddress(CompanyID)
	if err != nil {
		log.Println("Error from CH Find Registered Address call", err)
		q := url.Values{}
		q.Add("companyID", CompanyID)
		q.Add("error", "Cannot get Company Data from Companies House")
		redirectURL := findURL + "?" + q.Encode()
		log.Println("Redirecting to ", redirectURL)
		http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
		return
	}

	editPageVariables := roawa.NewEditPageVariables()
	editPageVariables.UpdateEditPageVariables(roaAddress, CompanyID)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	templateDirectory := cfg.TemplateDirectory

	templateFile := templateDirectory + "edit-registered-address.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Println("Error: Fail to parse Template :", templateFile, ":", err)
		panic(-1)
	}

	err = t.Execute(w, editPageVariables)
	if err != nil {
		log.Print("Template execution error: ", err)
	}
}
