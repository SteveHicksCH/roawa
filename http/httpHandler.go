package http

import (
	"log"
	"net/http"

	"github.com/shicks/roawa/config"
)

const urlPrefix = "/roawa"

// ServeHTTP facade for application HTTP Server
func ServeHTTP() {
	http.HandleFunc(urlPrefix+"/edit", editRegistedAddress)
	http.HandleFunc(urlPrefix+"/save", saveRegistedAddress)

	cfg, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}
	listenAddress := cfg.ListenAddress
	log.Println("Listen Address [", listenAddress, "]")

	err = http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Print("Fail to start HTTP Server: ", err)
		panic(-1)
	}
}
