package main

import (
	"html/template"
	"net/http"
	"path"
)

func (app *Application) render(w http.ResponseWriter, filname string, data interface{}) {

	fullPath := path.Join(app.templateDir, filname)
	tmpl, err := template.ParseFiles(fullPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
