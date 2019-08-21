package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/andaica/go-demo-api/article"
	"github.com/andaica/go-demo-api/db"
	"github.com/andaica/go-demo-api/user"

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
