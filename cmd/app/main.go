package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalln(err)
	}
}

func Root(wr http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Fprintf(wr, "<h1>SnippetBox</h1>")
}
