package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Data thing
type Data struct {
	Articles []Article
}

//Article Thing
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// HomePage thing
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from home")
}

// Articles thing
func Articles(w http.ResponseWriter, r *http.Request) {
	inits()
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(data)
}
func handleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/as", Articles)
	log.Fatal(http.ListenAndServe(":8082", nil))
}

var data = Data{}

func inits() {
	data.Articles = []Article{
		Article{Title: "title 1", Desc: "desc 1", Content: "test content"},
		Article{Title: "title 2", Desc: "desc 2", Content: "test content"},
	}
}
func main() {
	inits()
	fmt.Println(data)
	handleRequests()
}
