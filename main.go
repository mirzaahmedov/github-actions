package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	FirstName string
}

func getDataFromDatabase() (*User, error) {
	user := &User{}

	db, err := sql.Open("postgres", "postgres://test:test@localhost:5432/test?sslmode=disable")
	if err != nil {
		return nil, fmt.Errorf("Error connecting database, %v", err)
	}

	err = db.QueryRow(
		"SELECT first_name FROM users WHERE first_name='test'",
	).Scan(&user.FirstName)
	if err != nil {
		return nil, fmt.Errorf("Cannot get from database, %v", err)
	}

	return user, nil
}
