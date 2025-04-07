package main

import (
	"fmt"
	"lenslocked/src/views"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string, data any) error {
	t, err := views.Parse(filepath)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return err
	}
	return t.Execute(w, data)
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	err := executeTemplate(w, "templates/home.gohtml", map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	err := executeTemplate(w, "templates/contact.gohtml", map[string]any{
		"Name": name,
	})
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	err := executeTemplate(w, "templates/faq.gohtml", nil)
	if err != nil {
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
}

func main() {

	r := chi.NewRouter()

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
