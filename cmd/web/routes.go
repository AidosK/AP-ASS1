package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/student", app.students)
	mux.HandleFunc("/staff", app.staff)
	mux.HandleFunc("/applicant", app.applicant)
	mux.HandleFunc("/researcher", app.researcher)
	mux.HandleFunc("/contact", app.contact)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
