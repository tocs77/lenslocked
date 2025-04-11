package controllers

import (
	"lenslocked/src/views"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Static struct {
	Controllers struct {
		Home    func(w http.ResponseWriter, r *http.Request)
		Contact func(w http.ResponseWriter, r *http.Request)
		Faq     func(w http.ResponseWriter, r *http.Request)
	}
}

func (s *Static) Home(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	err := ExecuteTemplate(tpl, w, map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func (s *Static) Contact(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	err := ExecuteTemplate(tpl, w, map[string]any{
		"Name": name,
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func (s *Static) Faq(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	err := ExecuteTemplate(tpl, w, map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

type staticTemplates struct {
	Home    []string
	Contact []string
	Faq     []string
}

var staticTmpl staticTemplates = staticTemplates{
	Home:    []string{"tmpls/layout.gohtml", "tmpls/home.gohtml"},
	Contact: []string{"tmpls/layout.gohtml", "tmpls/contact.gohtml"},
	Faq:     []string{"tmpls/layout.gohtml", "tmpls/faq.gohtml"},
}

func (s *Static) SetupRoutes() {
	s.Controllers.Home = MakeHandler(s.Home, staticTmpl.Home...)
	s.Controllers.Contact = MakeHandler(s.Contact, staticTmpl.Contact...)
	s.Controllers.Faq = MakeHandler(s.Faq, staticTmpl.Faq...)
}
