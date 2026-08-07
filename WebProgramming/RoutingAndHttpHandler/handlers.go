package main

import (
	"fmt"
	"net/http"
)

var htmlContent = `
<!DOCTYPE html>
<html>
<head><title>%s</title></head>
<body>
	%s
	</body>
	</html>
`

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	homeContent := fmt.Sprintf(htmlContent, "Home", "<h1>Hello! welcome!</h1>")
	_, _ = w.Write([]byte(homeContent))
}

func (app *Application) about(w http.ResponseWriter, r *http.Request) {
	aboutContent := fmt.Sprintf(htmlContent, "About", "<h1>Just a software developer</h1>")
	_, _ = w.Write([]byte(aboutContent))
}

func (app *Application) contact(w http.ResponseWriter, r *http.Request) {
	contactContent := fmt.Sprintf(htmlContent, "Contact", "<h1>Contact me at @x : vexstack</h1>")
	_, _ = w.Write([]byte(contactContent))
}
