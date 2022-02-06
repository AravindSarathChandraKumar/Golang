package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var Articles []Article

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Description"`
	Content string `json:"Content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to HomePage!")
	fmt.Println("Success:HomePage")
}
func handleRequests() {
	//http.HandleFunc("/", homePage)
	//http.HandleFunc("/allarticles", returnAllArticles)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", getArticle)
	myRouter.HandleFunc("/article", createArticle).Methods("POST")
	http.Handle("/", myRouter)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing Endpoint:returnAllArticles")
	json.NewEncoder(w).Encode(Articles)

}

func getArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing Endpoint:returnSingleArticle")
	vars := mux.Vars(r)
	key := vars["id"]
	flag := 1
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
			flag = 1
			break
		} else {
			flag = 0
		}

	}
	if flag == 0 {
		fmt.Fprintf(w, "Requested Article Not Found!!!")
	}

}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)

}
func main() {
	Articles = []Article{
		{Id: "1", Title: "Hello1", Desc: "Article Description1", Content: "Article Content1"},
		{Id: "2", Title: "Hello2", Desc: "Article Description2", Content: "Article Content2"},
	}
	handleRequests()
}
