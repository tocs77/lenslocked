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

func main() {
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(os.Stdout, User{Name: "John", Meta: UserMeta{Visits: 10}})
}
