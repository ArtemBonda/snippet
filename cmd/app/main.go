package main

import (
	"log"
	"net"
	"net/http"

	"github.com/ArtemBonda/snippet/config"
)

func main() {
	cfg := config.NewConfig()

	mux := http.NewServeMux()
	mux.HandleFunc("/", Root)
	mux.HandleFunc("/snippet", ShowSnippet)
	mux.HandleFunc("/snippet/create", CreateSnippet)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	// go run cmd/app/main.go -port 9090
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalln(err)
	}
}
