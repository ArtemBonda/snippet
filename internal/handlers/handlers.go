package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//CreateSnippet хендлер для создания новой заметки
func (app *Application) CreateSnippet(wr http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wr.Header().Set("Allow", http.MethodPost)
		http.Error(wr, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	wr.Write([]byte("Create new notice..."))
}

//ShowSnippet вывод пользователю заметок
func (app *Application) ShowSnippet(wr http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(wr, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Fprintf(wr, "Snippet id = %d", id)
}

func (app *Application) Root(wr http.ResponseWriter, r *http.Request) {
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
		app.ErrorLog.Println(err.Error())
		http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(wr, nil)
	if err != nil {
		app.ErrorLog.Println(err.Error())
		http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
