package main

import (
	"net/http"
)

// var htmlContent = `
// <!DOCTYPE html>
// <html>
// <head><title>%s</title></head>
// <body>
// 	%s
// 	</body>
// 	</html>
// `

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	// homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>Hello! welcome!</h1>")
	// _, _ = w.Write([]byte(homeContent))
	//above is the old method

	//we gonna use render function now

	app.infoLog.Printf("Session data : %s", app.session.GetString(r, "userID"))
	app.render(w, r, "index.html", nil)

}

func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}

func (app *Application) contact(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "contact.html", nil)
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	app.session.Put(r, "userID", "123456")

	if r.Method == http.MethodPost{
		
	}

	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm)
	})
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.html", nil)
}
