package main

import (
	"net/http"
)

func (app *Application) render(w http.ResponseWriter, filname string, data interface{}) {

	// fullPath := path.Join(app.templateDir, filname)
	// tmpl, err := template.ParseFiles(fullPath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = tmpl.Execute(w, data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	//upgrading rendering with cache now

	if app.tp == nil {
		http.Error(w, "template rendering not set", http.StatusInternalServerError)
		return
	}
	app.tp.Render(w, filname, data)
}
