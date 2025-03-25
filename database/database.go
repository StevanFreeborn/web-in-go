package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, openErr := sql.Open("sqlite3", "./test.db")
	defer db.Close()

	if openErr != nil {
		fmt.Print(openErr)
	}

	connectionErr := db.Ping()

	if connectionErr != nil {
		fmt.Print(connectionErr)
	}

}
