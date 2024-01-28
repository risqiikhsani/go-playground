package inputOutput

import (
	"database/sql"
	"fmt"
)

func InputData(db *sql.DB) {
	fmt.Println("Enter name:")
	var name string
	fmt.Scanln(&name)

	fmt.Println("Enter email:")
	var email string
	fmt.Scanln(&email)

	// Insert data into the users table
	insertQuery := `
        INSERT INTO users (name, email) VALUES ($1, $2)
    `

	// Prepare the SQL statement
	stmt, err := db.Prepare(insertQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	// Execute the SQL statement with parameters
	_, err = stmt.Exec(name, email)
	if err != nil {
		panic(err)
	}

	fmt.Println("Data inserted successfully!")
}
