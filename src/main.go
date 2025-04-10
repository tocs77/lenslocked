package main

import (
	"fmt"
	"html/template"
	"lenslocked/src/controllers"
	"lenslocked/src/templates"
	"lenslocked/src/views"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(tpl views.Template, w http.ResponseWriter, data any) error {
	err := tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return err
	}
	return nil
}

func homeFunc(tpl views.Template, w http.ResponseWriter, _ *http.Request) {
	err := executeTemplate(tpl, w, map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func contactFunc(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	err := executeTemplate(tpl, w, map[string]any{
		"Name": name,
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func faqFunc(tpl views.Template, w http.ResponseWriter, _ *http.Request) {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{Question: "What is the capital of France?", Answer: "Paris"},
		{Question: "What is the capital of Germany?", Answer: "Berlin"},
		{Question: "What adress is google site?", Answer: "<a href=\"https://google.com\">Google</a>"},
		{Question: "What is the capital of Japan?", Answer: "Tokyo"},
	}
	err := executeTemplate(tpl, w, questions)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func signupFunc(tpl views.Template, w http.ResponseWriter, _ *http.Request) {
	err := executeTemplate(tpl, w, nil)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}
func main() {
	r := chi.NewRouter()
	homeHandler := controllers.StaticHandler(homeFunc, "tmpls/layout.gohtml", "tmpls/home.gohtml")
	contactHandler := controllers.StaticHandler(contactFunc, "tmpls/layout.gohtml", "tmpls/contact.gohtml")
	faqHandler := controllers.StaticHandler(faqFunc, "tmpls/layout.gohtml", "tmpls/faq.gohtml")
	signupHandler := controllers.StaticHandler(signupFunc, "tmpls/layout.gohtml", "tmpls/signup.gohtml")

	r.Get("/", homeHandler)
	r.Get("/contact/{name}", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/signup", signupHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	// Static files
	fileServer := http.FileServer(http.FS(templates.FSstatic))
	r.Handle("/static/*", fileServer)

	port := os.Getenv("APP_PORT")
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, r)
}
