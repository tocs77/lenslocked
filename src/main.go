package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		http.Error(w, "Page not found", http.StatusNotFound)
		return
	}
	err = t.Execute(w, map[string]any{
		"Date": time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		fmt.Println("Error executing template:", err)
		http.Error(w, "Page not found", http.StatusInternalServerError)
		return
	}
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	fmt.Fprintf(w, "<h1>Contact %s</h1>", name)
}

func faqHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, `
	<h1>FAQ</h1>
	<ul>
		<li>The first rule of Fight Club is: You do not talk about Fight Club.</li>
		<li>The second rule of Fight Club is: You do not talk about Fight Club.</li>
	</ul>
	`)
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
