package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTemplate *template.Template
}

func Parse(filepath string) (Template, error) {
	parsedTemplate, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("error parsing template: %w", err)
	}
	return Template{htmlTemplate: parsedTemplate}, nil
}

func (t *Template) Execute(w http.ResponseWriter, data any) error {
	w.Header().Set("Content-Type", "text/html")
	err := t.htmlTemplate.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return err
	}
	return nil
}

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
