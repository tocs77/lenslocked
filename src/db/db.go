package db

import (
	"database/sql"
	"fmt"
	"lenslocked/src/models"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var users = []models.NewUser{
	{Email: "john.doe@example.com", Name: "John Doe", Password: "password"},
	{Email: "jane.smith@example.com", Name: "Jane Smith", Password: "password"},
	{Email: "michael.johnson@example.com", Name: "Michael Johnson", Password: "password"},
	{Email: "emily.williams@example.com", Name: "Emily Williams", Password: "password"},
	{Email: "david.brown@example.com", Name: "David Brown", Password: "password"},
	{Email: "sarah.miller@example.com", Name: "Sarah Miller", Password: "password"},
	{Email: "robert.davis@example.com", Name: "Robert Davis", Password: "password"},
	{Email: "jennifer.wilson@example.com", Name: "Jennifer Wilson", Password: "password"},
	{Email: "william.taylor@example.com", Name: "William Taylor", Password: "password"},
	{Email: "olivia.anderson@example.com", Name: "Olivia Anderson", Password: "password"},
}

func Connect() (*sql.DB, error) {

	databaseUrl := "postgres://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("DB_NAME")
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return db, nil
}

func PrepareDb(db *sql.DB) {
	query, err := models.GetQuery("createTableUsers")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to get users query: %v\n", err)
		os.Exit(1)
	}
	_, err = db.Exec(query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create users table: %v\n", err)
		os.Exit(1)
	}

	// _, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (
	// 				id SERIAL PRIMARY KEY,
	// 				user_id INT NOT NULL,
	// 				amount INT NOT NULL,
	// 				description TEXT,
	// 				FOREIGN KEY (user_id) REFERENCES users(id)
	// 				)`)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "Unable to create table: %v\n", err)
	// 	os.Exit(1)
	// }
}

func FillUsers(us *models.UserService) {

	res, err := us.DB.Exec(`SELECT * FROM users`)
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
		for _, user := range users {
			us.Create(&user)
		}
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
		for _, user := range users {
			row := db.QueryRow(`INSERT INTO users (email, name, password_hash) VALUES ($1, $2, $3) RETURNING id`, user.Email, user.Name, user.Password)
			var id int
			scanErr := row.Scan(&id)
			if scanErr != nil {
				fmt.Fprintf(os.Stderr, "Unable to scan data: %v\n", scanErr)
				os.Exit(1)
			}
			fmt.Println("User id: ", id)
		}
	}
}
