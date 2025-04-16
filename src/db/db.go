package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() (*sql.DB, error) {

	var databaseUrl string
	databaseUrl = "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db, nil
}

func PrepareDb(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
					id SERIAL PRIMARY KEY, 
					email TEXT UNIQUE NOT NULL, 
					name TEXT
					)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
					id SERIAL PRIMARY KEY, 
					user_id INT NOT NULL, 
					amount INT NOT NULL, 
					description TEXT,
					FOREIGN KEY (user_id) REFERENCES users(id)
					)`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
		os.Exit(1)
	}
}

func FillDb(db *sql.DB) {
	res, err := db.Exec(`SELECT * FROM users`)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to select data: %v\n", err)
		os.Exit(1)
	}
	var rowsAffected int64
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get rows affected: %v\n", err)
		os.Exit(1)
	}
	if rowsAffected == 0 {
		_, err = db.Exec(`INSERT INTO users (email, name) VALUES ($1, $2)`, "test@test.com", "Test User")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to insert data: %v\n", err)
			os.Exit(1)
		}
	}
}
