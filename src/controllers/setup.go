package controllers

import (
	"lenslocked/src/templates"
	"lenslocked/src/views"
	"net/http"
)

func MakeHandler(handler func(views.Template, http.ResponseWriter, *http.Request), filepathes ...string) http.HandlerFunc {
	tpl := views.Must(views.ParseFS(templates.FS, filepathes...))
	return func(w http.ResponseWriter, r *http.Request) {
		handler(tpl, w, r)
	}
}

func ExecuteTemplate(tpl views.Template, w http.ResponseWriter, data any) error {
	err := tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return err
	}
	return nil
}
