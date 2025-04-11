package main

import (
	"fmt"
	"lenslocked/src/controllers"
	"lenslocked/src/templates"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	users := controllers.Users{}
	users.SetupRoutes()
	r.Get("/signup", users.Controllers.New)
	r.Post("/users", users.Controllers.Create)

	static := controllers.Static{}
	static.SetupRoutes()

	r.Get("/", static.Controllers.Home)
	r.Get("/contact/{name}", static.Controllers.Contact)
	r.Get("/faq", static.Controllers.Faq)

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
