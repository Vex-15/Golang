package main

import (
	"net/http"
)

func (app *Application) render(w http.ResponseWriter, r *http.Request, filname string, data *templateData) {

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

	mergedData := app.defaultTemplateData(data, r)
	app.tp.Render(w, filname, mergedData)

}

func (app *Application) defaultTemplateData(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}
	data.flash = app.session.PopString(r, "flash")
	return data
}
