package http

import "net/http"

// ServeHTTP facade for application HTTP Server
func ServeHTTP() {
	http.HandleFunc("/edit", editRegistedAddress)
	http.HandleFunc("/save", saveRegistedAddress)

	http.ListenAndServe("127.0.0.1:8080", nil)
}
