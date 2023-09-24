package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	FirstName string
}

func getDataFromDatabase() (*User, error) {
	user := &User{}

	db, err := sql.Open("postgres", "postgres://test:test@localhost:5432/test")
	if err != nil {
		return nil, err
	}

	err = db.QueryRow(
		"SELECT first_name FROM users WHERE first_name='test'",
	).Scan(&user.FirstName)
	if err != nil {
		return nil, err
	}

	return user, nil
}
