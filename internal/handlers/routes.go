package handlers

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", app.Root)
	mux.HandleFunc("/snippet", app.ShowSnippet)
	mux.HandleFunc("/snippet/create", app.CreateSnippet)

	fs := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	return mux
}