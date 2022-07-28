package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/ArtemBonda/snippet/config"
	"github.com/ArtemBonda/snippet/internal/handlers"
)

func main() {
	cfg := config.NewConfig()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Root)
	mux.HandleFunc("/snippet", handlers.ShowSnippet)
	mux.HandleFunc("/snippet/create", handlers.CreateSnippet)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	addr := net.JoinHostPort(cfg.Host, cfg.Port)
	server := &http.Server{
		Addr:     addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// go run cmd/app/main.go -port 9090
	infoLog.Printf("Starting server on %s", addr)
	if err := server.ListenAndServe(); err != nil {
		errorLog.Fatalln(err)
	}
}
