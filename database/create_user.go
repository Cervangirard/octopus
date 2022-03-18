package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var schema = `
CREATE TABLE IF NOT EXISTS users
(
    username INTEGER PRIMARY KEY AUTOINCREMENT,
    password TEXT,
    admin LOGICAL
)
`

func CreateUsers() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Printf("Cannot open database. err=%v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("Cannot check database connection. err=%v\n", err)
		os.Exit(1)
	}

	stmt, err := db.Prepare(schema)
	if err != nil {
		fmt.Printf("Cannot execute SQL query. err=%v\n", err)
		os.Exit(1)
	}

	stmt.Exec()

	fmt.Println("Ping to database successful!")
}
