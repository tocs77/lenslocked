package main

import (
	"fmt"
	"lenslocked/src/controllers"
	"lenslocked/src/db"
	"lenslocked/src/models"
	"lenslocked/src/templates"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	baseDb, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	db.PrepareDb(baseDb)
	defer baseDb.Close()

	userService := models.UserService{DB: baseDb}
	//db.FillUsers(&userService)
	//db.FillDb(baseDb)
	fmt.Println("Connected to database")

	r := chi.NewRouter()

	users := controllers.Users{UserService: &userService}
	users.SetupRoutes()
	r.Get("/signup", users.Controllers.New)
	r.Post("/users", users.Controllers.Create)
	r.Get("/signin", users.Controllers.SignIn)
	r.Post("/auth", users.Controllers.Auth)

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
