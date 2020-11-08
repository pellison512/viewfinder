package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pellison512/viewfinder/server/handlers/v2"
)

func main() {
	//http.HandleFunc("/hello", hello)
	r := mux.NewRouter()
	r.HandleFunc("/windows", handlers.PostWindowsHandler).Methods("POST")
	r.HandleFunc("/headers", handlers.HeadersHandler).Methods("GET")
	//	http.HandleFunc("/headers", handlers.HeadersHandler)
	//http.HandleFunc("/windows", handlers.WindowsHandler)
	http.ListenAndServe(":8090", nil)
}
