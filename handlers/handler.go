package Handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This Is Home Page ..."))
}

func ServeLogin(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "login.page.html")
}

func ServeSignup(w http.ResponseWriter, r *http.Request) {
	ServePages(w, "signup.page.html")

}

func ServePages(w http.ResponseWriter, tmpl string) {
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Println("Error Parsing template", err)
	}
}
