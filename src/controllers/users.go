package controllers

import (
	"lenslocked/src/views"
	"net/http"
)

type Users struct {
	Controllers struct {
		New func(w http.ResponseWriter, r *http.Request)
	}
}

func (u *Users) New(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(tpl, w, nil)
}

type userTemplates struct {
	New []string
}

var tml userTemplates = userTemplates{
	New: []string{"tmpls/layout.gohtml", "tmpls/signup.gohtml"},
}

func (u *Users) SetupRoutes() {
	u.Controllers.New = MakeHandler(u.New, tml.New...)
}
