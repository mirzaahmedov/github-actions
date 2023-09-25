package store

import "database/sql"

type Store struct {
	db *sql.DB
}

func New() *Store {
	return &Store{}
}
func (s *Store) Open() error {
	db, err := sql.Open("postgres", "postgres://postgres:@localhost:5432/godo?sslmode=disable")
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}
func (s *Store) Close() error {
	return s.db.Close()
}
