package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RouteHandler() {
	port := ":8080"
	r := mux.NewRouter()
	r.HandleFunc("/", LoginHandler)
	fmt.Println("Serving Port = ", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}

func main() {
	RouteHandler()
}
