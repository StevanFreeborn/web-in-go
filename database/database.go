package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, openErr := sql.Open("sqlite3", "./test.db")
	defer db.Close()

	if openErr != nil {
		fmt.Println(openErr)
	}

	connectionErr := db.Ping()

	if connectionErr != nil {
		fmt.Println(connectionErr)
	}

	createUsersTableQuery := `
    CREATE TABLE users (
      id INTEGER PRIMARY KEY AUTOINCREMENT,
      username TEXT NOT NULL,
      password TEXT NOT NULL,
      created_at DATETIME
    );
  `

	_, createUserTableError := db.Exec(createUsersTableQuery)

	if createUserTableError != nil {
		fmt.Println(createUserTableError)
	}

	newUserUsername := "john.wick"
	newUserPassword := "supersecretpassword"
	newUserCreatedAt := time.Now()

	createUserQuery := `
    INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)
  `

	createUserResult, createUserErr := db.Exec(
		createUserQuery,
		newUserUsername,
		newUserPassword,
		newUserCreatedAt,
	)

	if createUserErr != nil {
		fmt.Println(createUserErr)
	}

	userId, userIdErr := createUserResult.LastInsertId()

	if userIdErr != nil {
		fmt.Println(userIdErr)
	}

	fmt.Println(userId)

	var username string

	getUserQuery := "SELECT username FROM users WHERE id = ?"
	getUserErr := db.QueryRow(getUserQuery, userId).Scan(&username)

	if getUserErr != nil {
		fmt.Println(getUserErr)
	}

	fmt.Println(username)

	userRows, getUsersErr := db.Query("SELECT * FROM users")

	if getUsersErr != nil {
		fmt.Println(getUsersErr)
	}

	defer userRows.Close()

	type user struct {
		id        int
		username  string
		password  string
		createdAt time.Time
	}

	var users []user

	for userRows.Next() {
		var u user

		getUserErr := userRows.Scan(&u.id, &u.username, &u.password, &u.createdAt)

		if getUserErr != nil {
			fmt.Println(getUserErr)
		}

		users = append(users, u)
	}

	userRowsErr := userRows.Err()

	if userRowsErr != nil {
		fmt.Println(userRowsErr)
	}

	fmt.Printf("%#v", users)
}
