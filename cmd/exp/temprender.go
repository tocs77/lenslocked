package main

import (
	"html/template"
	"log"
	"os"
)

type User struct {
	Name string
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

// RenderTemplate renders a template with user data
func RenderTemplate() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, User{Name: "John", Meta: UserMeta{Visits: 10}})
}
