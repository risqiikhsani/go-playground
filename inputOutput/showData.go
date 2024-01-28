package inputOutput

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ShowData(db *sql.DB) {
	selectQuery := `
	SELECT id, name, email FROM users`

	// Execute the SELECT query
	rows, err := db.Query(selectQuery)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Iterate over the rows and print the data
	fmt.Println("User Data:")
	for rows.Next() {
		var id int
		var name, email string
		if err := rows.Scan(&id, &name, &email); err != nil {
			panic(err)
		}
		fmt.Printf("ID: %d, Name: %s, Email: %s\n", id, name, email)
	}
	if err := rows.Err(); err != nil {
		panic(err)
	}

}
