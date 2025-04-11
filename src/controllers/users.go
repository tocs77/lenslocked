package controllers

import (
	"lenslocked/src/views"
	"net/http"
)

type Users struct {
	Controllers struct {
		New    func(w http.ResponseWriter, r *http.Request)
		Create func(w http.ResponseWriter, r *http.Request)
	}
}

func (u *Users) New(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(tpl, w, nil)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var email string
	var password string

	email = r.FormValue("email")
	password = r.FormValue("password")

	w.Write([]byte("Create user " + email + " with password " + password))
}

type userTemplates struct {
	New []string
}

var tmpl userTemplates = userTemplates{
	New: []string{"tmpls/layout.gohtml", "tmpls/signup.gohtml"},
}

func (u *Users) SetupRoutes() {
	u.Controllers.New = MakeHandler(u.New, tmpl.New...)
	u.Controllers.Create = u.Create
}
