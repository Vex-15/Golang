package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicPath))))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/login", app.login)
	mux.HandleFunc("/register", app.register)
	mux.HandleFunc("/contact", app.contact)

	handler := app.recover(app.logger(app.session.Enable(mux)))
	return handler
}
