package http

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/shicks/roawa"
	"github.com/shicks/roawa/config"
	"github.com/shicks/roawa/restclient"
)

func saveRegistedAddress(w http.ResponseWriter, r *http.Request) {
	confirmPageVars := roawa.NewConfirmPageVariables()
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(30000)
	confirmPageVars.Reference = strconv.Itoa(random)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	templateDirectory := cfg.TemplateDirectory
	templateFile := templateDirectory + "confirm-save.html"
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Print("Error - Fail to parse Template:", templateFile, ":", err)
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
	companyID := r.FormValue("companyID")

	//  Do save via ROA API (add a new failure page to display on errors)

	transactionRequest := roawa.TransactionRequest{
		CompanyNumber: companyID,
		Reference:     confirmPageVars.Reference,
		Description:   "ROAWA Update",
	}
	transactionResponse, err := restclient.CreateTransaction(transactionRequest)
	if err != nil {
		confirmPageVars.Error = "Fail setting up change request"
		log.Print("Error - Fail create a new transaction:", confirmPageVars, err)
		executeTemplate(t, w, confirmPageVars)
		return
	}
	if transactionResponse.ID == "" {
		confirmPageVars.Error = "Fail setting up change request"
		log.Print("No Transaction id returned in response - ", transactionResponse)
		executeTemplate(t, w, confirmPageVars)
		return
	}

	roaAddress := roawa.RoaAddress{}
	roaAddress.UpdateRoaAddress(confirmPageVars)
	err = restclient.ReplaceROA(roaAddress, transactionResponse.ID)
	if err != nil {
		confirmPageVars.Error = "Fail updating the ROA"
		log.Print("Error - Fail in ROA Update:", templateFile, ":", err)
		executeTemplate(t, w, confirmPageVars)
		return
	}

	executeTemplate(t, w, confirmPageVars)

}

func executeTemplate(t *template.Template, w http.ResponseWriter, confirmPageVars roawa.ConfirmPageVariables) {
	err := t.Execute(w, confirmPageVars)
	if err != nil {
		log.Print("Template execution error: ", err)
	}
}
