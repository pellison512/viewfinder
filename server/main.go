package main

import (
	"net/http"

	"github.com/pellison512/viewfinder/server/handlers/v2"
)

func main() {
	//http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", handlers.Headers)
	http.ListenAndServe(":8090", nil)
}
