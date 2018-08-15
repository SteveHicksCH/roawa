package http

import (
	"html/template"
	"log"
	"net/http"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
)

func findCompany(w http.ResponseWriter, r *http.Request) {
	findPageVars := roawa.NewFindPageVariables()

	log.Println("find company landing page")

	// Look for query parameters in  error scenario
	keys, ok := r.URL.Query()["error"]
	if ok && len(keys[0]) > 0 {
		findPageVars.Error = keys[0]
	}
	keys, ok = r.URL.Query()["companyID"]
	if ok && len(keys[0]) > 0 {
		findPageVars.CompanyID = keys[0]
	}

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	templateDirectory := cfg.TemplateDirectory
	templateFile := templateDirectory + "find-company.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Print("Error - Fail to parse Template:", templateFile, ":", err)
		panic(-1)
	}

	err = t.Execute(w, findPageVars)
	if err != nil {
		log.Print("Template execution error: ", err)
	}

}
