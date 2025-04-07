package main

import (
	"fmt"
	"lenslocked/src/controllers"
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
	err := executeTemplate(tpl, w, nil)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func main() {
	r := chi.NewRouter()
	homeHandler := controllers.StaticHandler("templates/home.gohtml", homeFunc)
	contactHandler := controllers.StaticHandler("templates/contact.gohtml", contactFunc)
	faqHandler := controllers.StaticHandler("templates/faq.gohtml", faqFunc)

	r.Get("/", homeHandler)
	r.Get("/contact/{name}", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	port := os.Getenv("APP_PORT")
	fmt.Println("Starting server on port " + port)
	http.ListenAndServe(":"+port, r)
}
