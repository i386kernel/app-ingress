package main

import (
	"fmt"
	"net/http"
	"os"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	servefile := currentDir + "/static/login.html"
	fmt.Println(servefile)
	http.ServeFile(w, r, servefile)
}
