package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Article struct {
	Judul   string `json:judul`
	Penulis string `json:penulis`
}

type Articles []Article

func main() {
	http.HandleFunc("/", getHome)
	http.HandleFunc("/articles", getArticles)
	http.HandleFunc("/post-article", postArticle)
	log.Printf("Service run in localhost:3000")
	http.ListenAndServe(":3000", nil)

}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Halaman Home"))
}

func getArticles(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET:: /articles")
	article := Articles{
		Article{Judul: "Test Judul", Penulis: "Heru"},
	}

	json.NewEncoder(w).Encode(article)
}

func postArticle(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST:: /post-article")
	if r.Method == "POST" {
		// body, err := ioutil.ReadAll(r.Body)

		// if err != nil {
		// 	log.Printf("Can't Read Body")
		// 	http.Error(w, "Can't Read Body", http.StatusBadRequest)
		// }

		// w.Write([]byte(string(body)))

		var article Article
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			log.Printf("Bad Request", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		if article.Judul == "" || article.Penulis == "" {
			log.Printf("Judul atau Penulis tidak boleh kosong")
			http.Error(w, "Bad Request, Judul atau Penulis tidak boleh kosong", http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(article)

	} else {
		log.Printf("invalid method")
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
}
