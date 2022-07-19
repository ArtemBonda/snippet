package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Fatalln(err)
	}
}

//CreateSnippet хендлер для создания новой заметки
func CreateSnippet(wr http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wr.Header().Set("Allow", http.MethodPost)
		http.Error(wr, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	wr.Write([]byte("Create new notice..."))
}

//ShowSnippet вывод пользователю заметок
func ShowSnippet(wr http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Fprintf(wr, "Snippet id = %d", id)
}

func Root(wr http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(wr, nil)
	if err != nil {
		http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
