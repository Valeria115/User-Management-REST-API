package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := connectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	code := m.Run()

	db.Exec("DELETE FROM users")

	os.Exit(code)
}

func setupTestData() {
	db.Exec("DELETE FROM users")
	db.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 3")

	db.Exec(
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
		1, "Alice", "alice@example.com",
	)
	db.Exec(
		"INSERT INTO users (id, name, email) VALUES ($1, $2, $3)",
		2, "Bob", "bob@example.com",
	)
}
