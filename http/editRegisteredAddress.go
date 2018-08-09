package http

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
	"github.com/shicks/roawa/restclient"
)

const testCompanyNumber = "00000006"

// editRegistedAddress is a handler for when a user starts to edit the company registered address.EditRegistedAddress.EditRegistedAddress
// It populates the Edit Registered Address Screen
func editRegistedAddress(w http.ResponseWriter, r *http.Request) {

	roaAddress, _ := restclient.FindRegisteredAddress(testCompanyNumber)

	editPageVariables := roawa.NewEditPageVariables()
	editPageVariables.UpdateEditPageVariables(roaAddress)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	templateDirectory := cfg.TemplateDirectory

	t, err := template.ParseFiles(templateDirectory + "edit-registered-address.html")
	if err != nil {
		log.Print("Fail to parse Template: ", err)
		panic(-1)
	}

	err = t.Execute(w, editPageVariables)
	if err != nil {
		log.Print("Template execution error: ", err)
	}
}
