package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	// Create users table if it doesn't exist
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		return nil, fmt.Errorf("error creating table: %v", err)
	}

	// Insert some sample data if the table is empty
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("error checking table: %v", err)
	}

	if count == 0 {
		insertData := `
		INSERT INTO users (name, email) VALUES
		('Alice Smith', 'alice@example.com'),
		('Bob Jones', 'bob@example.com'),
		('Carol White', 'carol@example.com');`

		_, err = db.Exec(insertData)
		if err != nil {
			return nil, fmt.Errorf("error inserting sample data: %v", err)
		}
	}

	return db, nil
}

func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, fmt.Errorf("error scanning user: %v", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %v", err)
	}

	return users, nil
}
