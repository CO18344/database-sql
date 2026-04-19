package main 

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main(){
	// Connect to the Cockroach database
	var db *sql.DB;
	var err error;
	db, err = sql.Open("postgres","postgresql://evermile:password@localhost:26257/users?sslmode=verify-full&sslrootcert=/home/satvik/.cockroach-certs/ca.crt")

	// Check if connection failed
	if err != nil {
		fmt.Println("Failed to connect to database")
		// fmt.Println(err)
		return
	}

	// Ensure database connection is closed when main exits
	defer db.Close()

	// Create a background context for database operations
	var ctx context.Context;
	ctx = context.Background();

	// Begin transaction -> gives *sql.Tx
	var tx *sql.Tx;
	tx, _ = db.BeginTx(ctx, nil)

	// Run SQL inside transaction
	var firstQueryErr error;
	_, firstQueryErr = tx.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", 1, "Satvik", "satviksingh600@gmail.com")

	// Check if first insert failed
	if(firstQueryErr != nil){
		fmt.Println("Error while running first query")
		fmt.Println(firstQueryErr)
	}

	// Execute second INSERT statement and capture error
	var secQueryErr error;
	_ ,secQueryErr = tx.Exec("INSERT INTO users (id, name, email) VALUES ($1, $2, $3)", 2, "Mohit", "mohitsharma@gmail.com")

	// Check if second insert failed
	if(secQueryErr != nil){
		fmt.Println("Error while running second query")
		fmt.Println(secQueryErr)
	}

	// Commit the transaction to save all changes
	tx.Commit()

	// Query and retrieve user data from the database
	var name string;
	db.QueryRow("SELECT name FROM users WHERE id = $1", 1).Scan(&name)

	fmt.Println(name)
}