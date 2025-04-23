package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	Email        string `json:"email"`
	Name         string `json:"name"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Email    string
	Name     string
	Password string
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(user *NewUser) (*User, error) {
	email := strings.ToLower(user.Email)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	var createdUser = User{
		Email:        email,
		Name:         user.Name,
		PasswordHash: string(hashedPassword),
	}

	query, err := GetQuery("createUser")
	if err != nil {
		return nil, err
	}

	row := us.DB.QueryRow(query, email, user.Name, string(hashedPassword))
	err = row.Scan(&createdUser.ID)
	if err != nil {
		return nil, err
	}

	return &createdUser, nil
}

func (us *UserService) Authenticate(email, password string) (*User, error) {
	email = strings.ToLower(email)
	query, err := GetQuery("findUserByEmail")
	if err != nil {
		return nil, err
	}

	row := us.DB.QueryRow(query, email)
	var user User
	err = row.Scan(&user.ID, &user.Email, &user.Name, &user.PasswordHash)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, err
	}
	fmt.Println("User authenticated")
	return &user, nil
}
