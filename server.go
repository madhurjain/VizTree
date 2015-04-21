package main

import (
	"./autocomplete"
	"bufio"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Cache templates
var templates = template.Must(template.ParseFiles("home.html"))

var trie *autocomplete.Trie = autocomplete.NewTrie()

type Response struct {
	Words []string `json:"words"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// 404 for all other url path
	if r.URL.Path[1:] != "" {
		http.NotFound(w, r)
		return
	}
	// dev
	// t, _ := template.ParseFiles("home.html")
	// t.Execute(w, nil)
	templates.ExecuteTemplate(w, "home.html", nil)
}

func suggestHandler(w http.ResponseWriter, r *http.Request) {
	// log.Println("-- suggestHandler")
	r.ParseForm()
	prefix := r.FormValue("prefix")
	words := trie.SearchWords(prefix)
	response := Response{Words: words}
	var toWrite []byte
	toWrite, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(toWrite)
}

func LoadTrie() {
	file, err := os.Open("linuxwords.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trie.AddWord(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Add all the words in dictionary to Trie
	LoadTrie()

	var httpHost string = os.Getenv("HOST")
	var httpPort string = os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/suggest/", suggestHandler)

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	log.Printf("server listening on %s:%s\n", httpHost, httpPort)
	http.ListenAndServe(httpHost+":"+httpPort, nil)
}
