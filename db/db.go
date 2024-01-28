package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Db() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "myuser"
		password = "mypassword"
		dbname   = "mydatabase"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Check if the connection to the database is successful
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to the database!")

	// Perform database operations here...

	// Create table statement
	createTableQuery := `
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            name VARCHAR(100),
            email VARCHAR(100)
        )
    `

	// Execute the create table statement
	_, err = db.Exec(createTableQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("Table created successfully!")

	return db
}
