package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RouteHandler() {
	port := ":80"
	r := mux.NewRouter()
	r.HandleFunc("/v3", LoginHandler)
	fmt.Println("Serving Port = ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	RouteHandler()
}
