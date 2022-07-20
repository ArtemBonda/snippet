package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"strconv"
)

func main() {
	host := flag.String("host", "localhost", "usage <localhost>'")
	port := flag.String("port", ":8080", "usage :<port>")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	addr := net.JoinHostPort(*host, *port)
	log.Printf("Запуск сервера на %s:%s\n", *host, *port)
	// go run cmd/app/main.go -port 9090
	if err := http.ListenAndServe(addr, mux); err != nil {
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
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
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
