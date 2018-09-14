package http

import (
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/justinas/alice"

	"github.com/shicks/roawa/config"
)

const urlPrefix = "/roawa"
const findURL = urlPrefix + "/find"

// ServeHTTP facade for application HTTP Server
func ServeHTTP() {

	router := pat.New()
	router.Get(findURL, findCompany)
	router.Get(urlPrefix+"/edit", editRegistedAddress)
	router.Post(urlPrefix+"/save", saveRegistedAddress)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	listenAddress := cfg.ListenAddress
	log.Println("Listen Address [", listenAddress, "]")

	chain := alice.New(handlerLogger)
	err = http.ListenAndServe(listenAddress, chain.Then(router))
	if err != nil {
		log.Print("Fail to start HTTP Server: ", err)
		panic(-1)
	}
}
