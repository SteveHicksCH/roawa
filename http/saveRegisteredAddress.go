package http

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
)

func saveRegistedAddress(w http.ResponseWriter, r *http.Request) {
	confirmPageVars := roawa.NewConfirmPageVariables()
	confirmPageVars.Reference = "10989097"

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	templateDirectory := cfg.TemplateDirectory
	templateFile := templateDirectory + "confirm-save.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Print("Fail to parse Template: ", err)
		panic(-1)
	}

	confirmPageVars.Premises = r.FormValue("premises")
	confirmPageVars.AddressLine1 = r.FormValue("addressLine1")
	confirmPageVars.AddressLine2 = r.FormValue("addressLine2")
	confirmPageVars.Town = r.FormValue("town")
	confirmPageVars.County = r.FormValue("county")
	country := r.FormValue("country")
	confirmPageVars.Country = confirmPageVars.Countries[country]
	confirmPageVars.Postcode = r.FormValue("postcode")
	confirmPageVars.POBox = r.FormValue("pOBox")

	//  Do save via ROA API

	err = t.Execute(w, confirmPageVars)
	if err != nil {
		log.Print("Template execution error: ", err)
	}

}
