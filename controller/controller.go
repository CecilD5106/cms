package controller

import (
	"html/template"
	"net/http"
)

var (
	homeController home
)

// Startup is the function that defines the templates
func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	homeController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}