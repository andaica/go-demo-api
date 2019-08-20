package main

import (
	"demoAPI/article"
	"demoAPI/db"
	"demoAPI/user"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequest() {
	myRouter := mux.NewRouter()

	user.RegistRouter(myRouter)
	article.RegistRouter(myRouter)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("REST API v2.0 - Mux Routers")
	db.Connect()
	handleRequest()
}
