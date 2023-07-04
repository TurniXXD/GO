package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/TurniXXD/GO/env"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{
			Title:   "Post 1",
			Desc:    "First post",
			Content: "Fuck you all",
		},
	}
	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

// capitalize letters to make function publically available
func HandleRequests() {
	// Higher order function
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	println("\nserving on port " + env.Process("GO_API_PORT"))
	log.Fatal(http.ListenAndServe(":"+env.Process("GO_API_PORT"), nil))
}
