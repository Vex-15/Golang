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
	app.render(w, "index.html", nil)

}

func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	// aboutContent := fmt.Sprintf(htmlContent, "About", "<h1>Just a software developer</h1>")
	// _, _ = w.Write([]byte(aboutContent))
	app.render(w, "about.html", nil)
}

func (app *Application) contact(w http.ResponseWriter, r *http.Request) {
	// contactContent := fmt.Sprintf(htmlContent, "Contact", "<h1>Contact me at @x : vexstack</h1>")
	// _, _ = w.Write([]byte(contactContent))
	app.render(w, "contact.html", nil)
}

func (app *Application) login(w http.ResponseWriter, r *http.Request) {
	app.session.Put(r, "userID", "123456")
	app.render(w, "login.html", nil)
}

func (app *Application) register(w http.ResponseWriter, r *http.Request) {
	app.render(w, "register.html", nil)
}
