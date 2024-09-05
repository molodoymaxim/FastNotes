package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	// Используем методы из структуры в качестве обработчиков маршрутов.
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/note", app.showSnippet)
	mux.HandleFunc("/note/create", app.createSnippet)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
