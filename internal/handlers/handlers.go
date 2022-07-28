package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strconv"
)

//CreateSnippet хендлер для создания новой заметки
func (app *Application) CreateSnippet(wr http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		wr.Header().Set("Allow", http.MethodPost)
		app.ClientError(wr, http.StatusMethodNotAllowed)
		return
	}
	wr.Write([]byte("Create new notice..."))
}

//ShowSnippet вывод пользователю заметок
func (app *Application) ShowSnippet(wr http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.NotFound(wr)
		return
	}
	fmt.Fprintf(wr, "Snippet id = %d", id)
}

func (app *Application) Root(wr http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(wr)
		return
	}
	files := []string{
		"./ui/html/home.page.html",
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(wr, err)
		return
	}

	err = tmpl.Execute(wr, nil)
	if err != nil {
		app.ServerError(wr, err)
		return
	}
}

//ServerError записывает сообщение об ошибке в errorLog, пользователю будет в ответе возвращен код ошибки
func (app *Application) ServerError(wr http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)

	http.Error(wr, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

//ClientError отправляет определенный код состояния и соответсвующее описание пользвателю.
func (app *Application) ClientError(wr http.ResponseWriter, status int) {
	http.Error(wr, http.StatusText(status), status)
}

//NotFound отправляет пользователю в ответе код 404
func (app *Application) NotFound(wr http.ResponseWriter) {
	app.ClientError(wr, http.StatusNotFound)
}
