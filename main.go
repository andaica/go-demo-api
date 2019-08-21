package main

import (
	"fmt"
	"log"
	"net/http"

	"go-demo-api/article"
	"go-demo-api/db"
	"go-demo-api/user"

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
