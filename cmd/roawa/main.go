package main

import (
	"log"

	"github.com/shicks/roawa"

	"github.com/shicks/roawa/config"
	"github.com/shicks/roawa/http"
)

func main() {

	initialise()
	http.ServeHTTP()
}

func initialise() {
	_, err := config.Get()
	if err != nil {
		log.Println("Error getting config:", err)
		panic(1)
	}

	c := *(roawa.GetCountries())
	log.Println(len(c), " entries in countries map")
}
