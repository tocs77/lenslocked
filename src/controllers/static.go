package controllers

import (
	"lenslocked/src/views"
	"net/http"
)

func StaticHandler(filepath string, handler func(views.Template, http.ResponseWriter, *http.Request)) (http.HandlerFunc, error) {
	tpl, err := views.Parse(filepath)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		handler(tpl, w, r)
	}, nil
}
