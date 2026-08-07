package main

import "net/http"

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("./", app.home)
	mux.HandleFunc("./about", app.about)
	mux.HandleFunc("./contact", app.contact)
	return mux
}
