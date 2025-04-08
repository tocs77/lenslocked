package controllers

import (
	"lenslocked/src/templates"
	"lenslocked/src/views"
	"net/http"
)

func StaticHandler(handler func(views.Template, http.ResponseWriter, *http.Request), filepathes ...string) http.HandlerFunc {
	tpl := views.Must(views.ParseFS(templates.FS, filepathes...))
	return func(w http.ResponseWriter, r *http.Request) {
		handler(tpl, w, r)
	}
}
