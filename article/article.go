package article

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/andaica/go-demo-api/authen"

	"github.com/gorilla/mux"
)

func RegistRouter(router *mux.Router) {
	router.HandleFunc("/api/articles", authen.BasicAuth(createArticle)).Methods("POST")
	router.HandleFunc("/api/articles", getAllArticles)
}

var dama DataMapping

func createArticle(w http.ResponseWriter, r *http.Request) {
	article := getDataFromPost(r)
	log.Println("Endpoint hit: createArticle ", article)
	userId := authen.GetAuthenticatedUserId(r)
	result, isOk := dama.insertNewArticle(article, userId)
	if isOk {
		responseSuccess(w, result)
	} else {
		responseError(w, "Can not create Article!")
	}
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: AllArticles")
	articles := dama.fetchAll()
	json.NewEncoder(w).Encode(articles)
}

func getDataFromPost(r *http.Request) Article {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req Request
	json.Unmarshal(reqBody, &req)
	return req.Article
}

func responseSuccess(w http.ResponseWriter, article ArticleResponse) {
	response := Response{article, "OK"}
	json.NewEncoder(w).Encode(response)
}

func responseError(w http.ResponseWriter, message string) {
	response := map[string]interface{}{"status": "NG", "message": message}
	json.NewEncoder(w).Encode(response)
}
