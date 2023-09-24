package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	got, err := getDataFromDatabase()
	if err != nil {
		t.Fatal("get user from database failed")
	}

	if got.FirstName != "test" {
		t.Fatal("got incorrect user from database")
	}
}
