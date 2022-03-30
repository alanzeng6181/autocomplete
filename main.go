package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alanzeng6181/autocomplete/datastructure"
)

func search(w http.ResponseWriter, r *http.Request) {
	keywords, ok := r.URL.Query()["keyword"]
	if !ok || len(keywords) < 1 {
		w.WriteHeader(400)
		w.Write([]byte("keyword is not specified"))
	}

	topN := trie.GetTopN(datastructure.MakeSearchString1(keywords[0]))
	arr := make([]string, 0)
	for _, s := range topN {
		arr = append(arr, s.Text)
	}
	data, _ := json.Marshal(arr)
	w.Write(data)
}

var trie *datastructure.Trie[datastructure.SearchString] = datastructure.NewTrie[datastructure.SearchString]()

func add(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" && r.Method != "PUT" {
		w.WriteHeader(405)
		return
	}

	keywords, ok := r.URL.Query()["keyword"]
	if !ok || len(keywords) < 1 {
		w.WriteHeader(400)
		w.Write([]byte("keyword is not specified"))
	}

	_, err := trie.Add(datastructure.MakeSearchString(keywords[0], 1))

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error occured while adding keyword %s, due to %v", keywords[0], err)))
	} else {
		w.WriteHeader(200)
	}
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/add", add)
	http.HandleFunc("/search", search)
	const port int32 = 8080
	log.Printf("listening at %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
