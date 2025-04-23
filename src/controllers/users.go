package controllers

import (
	"encoding/json"
	"lenslocked/src/models"
	"lenslocked/src/views"
	"net/http"
)

type Users struct {
	Controllers struct {
		New    func(w http.ResponseWriter, r *http.Request)
		Create func(w http.ResponseWriter, r *http.Request)
		SignIn func(w http.ResponseWriter, r *http.Request)
		Auth   func(w http.ResponseWriter, r *http.Request)
	}
	UserService *models.UserService
}

func (u *Users) New(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(tpl, w, nil)
}

func (u *Users) SignIn(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	ExecuteTemplate(tpl, w, nil)
}

func (u *Users) Auth(w http.ResponseWriter, r *http.Request) {
	var email string
	var password string

	email = r.FormValue("email")
	password = r.FormValue("password")
	user, err := u.UserService.Authenticate(email, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the response as JSON
	json.NewEncoder(w).Encode(user)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var email string
	var password string
	var name string
	email = r.FormValue("email")
	name = r.FormValue("name")
	password = r.FormValue("password")
	user, err := u.UserService.Create(&models.NewUser{Email: email, Name: name, Password: password})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Create user " + user.Email + " with password " + user.PasswordHash))
}

type userTemplates struct {
	New    []string
	SignIn []string
}

var tmpl userTemplates = userTemplates{
	New:    []string{"tmpls/layout.gohtml", "tmpls/signup.gohtml"},
	SignIn: []string{"tmpls/layout.gohtml", "tmpls/signin.gohtml"},
}

func (u *Users) SetupRoutes() {
	u.Controllers.New = MakeHandler(u.New, tmpl.New...)
	u.Controllers.Create = u.Create
	u.Controllers.SignIn = MakeHandler(u.SignIn, tmpl.SignIn...)
	u.Controllers.Auth = u.Auth
}
