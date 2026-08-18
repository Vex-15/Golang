package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *Application) routes() http.Handler {
	mux := http.NewServeMux()

	defaultMiddleware := alice.New(app.recover, app.logger)
	secureMiddleware := defaultMiddleware.Append(app.session.Enable)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicPath))))
	mux.HandleFunc("/", secureMiddleware.ThenFunc(app.home))
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/login", secureMiddleware.ThenFunc(app.login))
	mux.HandleFunc("/register", secureMiddleware.ThenFunc(app.register))
	mux.HandleFunc("/contact", app.contact)
	mux.HandleFunc("/submit", secureMiddleware.Append(app.requireAuth).ThenFunc(app.submit))

	handler := secureMiddleware.Then(mux)
	return handler
}
