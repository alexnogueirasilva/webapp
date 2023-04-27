package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates sets the templates variable with all html files
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate executes a template
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	err := templates.ExecuteTemplate(w, template, data)
	if err != nil {
		return
	}
}
