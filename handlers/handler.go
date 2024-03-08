package Handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ServeHomePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home page ..."))
}

func ServeHome(w http.ResponseWriter, tmpl string) {
	parsedTemplates, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplates.Execute(w, nil)
	if err != nil {
		log.Println("Error Parsing template", err)
	}

}
