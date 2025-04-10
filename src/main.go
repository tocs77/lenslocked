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

func homeFunc(tpl views.Template, w http.ResponseWriter, _ *http.Request) {
	err := controllers.ExecuteTemplate(tpl, w, map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func contactFunc(tpl views.Template, w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	err := controllers.ExecuteTemplate(tpl, w, map[string]any{
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
	err := controllers.ExecuteTemplate(tpl, w, questions)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func signupFunc(tpl views.Template, w http.ResponseWriter, _ *http.Request) {
	err := controllers.ExecuteTemplate(tpl, w, nil)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}
func main() {
	r := chi.NewRouter()
	homeHandler := controllers.MakeHandler(homeFunc, "tmpls/layout.gohtml", "tmpls/home.gohtml")
	contactHandler := controllers.MakeHandler(contactFunc, "tmpls/layout.gohtml", "tmpls/contact.gohtml")
	faqHandler := controllers.MakeHandler(faqFunc, "tmpls/layout.gohtml", "tmpls/faq.gohtml")

	users := controllers.Users{}
	users.SetupRoutes()

	r.Get("/", homeHandler)
	r.Get("/contact/{name}", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/signup", users.Controllers.New)
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
