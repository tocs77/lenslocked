package main

import (
	"errors"
	"fmt"
)

func StartConnection() {
	err := CreateUser()
	if err != nil {
		fmt.Println(err)
	}
}

func Connect() error {
	return errors.New("connection error")
}

func CreateUser() error {
	err := Connect()
	if err != nil {
		return fmt.Errorf("create user error: %w", err)
	}
	return nil
}
