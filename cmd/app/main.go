package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ArtemBonda/snippet/config"
	"github.com/ArtemBonda/snippet/internal/handlers"
)

func main() {
	cfg := config.NewConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	// go run cmd/app/main.go -port 9090
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalln(err)
	}
}
