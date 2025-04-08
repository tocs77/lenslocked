package controllers

import (
	"lenslocked/src/templates"
	"lenslocked/src/views"
	"net/http"
)

func StaticHandler(filepath string, handler func(views.Template, http.ResponseWriter, *http.Request)) http.HandlerFunc {
	tpl := views.Must(views.ParseFS(templates.FS, filepath))
	return func(w http.ResponseWriter, r *http.Request) {
		handler(tpl, w, r)
	}
}
