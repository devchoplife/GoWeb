package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Installing the go-sql-driver/mysql package
//In Go, database driver is a package which implements the low level
//details of a specific database (in our case MySQL)

//go get -u github.com/go-sql-driver/mysql

//Connecting to a MySQL database
//first import the required packages
/*
import "database/sql"
import _ "go-sql-driver/mysql"
*/

//Then use the username and password of the database to create a conection like below:
//db, err := sql.Open("mysql", "username:password@(127.0.0.1:3306)/dbname?parseTime=true")

func main() {
	//Create a connection to the database
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/root?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	{ //Create a new table
		query := `
		CREATE TABLE users (
			id INT AUTO_INCREMENT,
			username TEXT NOT NULL,
			password TEXT NOT NULL,
			created_at DATETIME,
			PRIMARY KEY (id)
		);`

		//Execute query and check for errors
		if _, err := db.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		//Insert a new user
		username := "johndoe"
		password := "secret"
		createdAt := time.Now()

		result, err := db.Exec(`INSERT INTO users (usernsme, password, createdAt) VALUES (?, ?, ?)`, username, password, createdAt)
		if err != nil {
			log.Fatal(err)
		}

		id, err := result.LastInsertId()
		fmt.Println(id)
	}

	{
		//Querying our database

		//Note: db.Query is used to query multiple rows
		//db.QueryRow is used to query a specific row

		//Query a single user
		var (
			id        int
			username  string
			password  string
			createdAt time.Time
		)

		query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
		if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, username, password, createdAt)
	}

	{
		//Query all users
		type user struct {
			id        int
			username  string
			password  string
			createdAt time.Time
		}

		rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var users []user
		for rows.Next() {
			var u user

			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%#v", users)
	}

	{
		//delete user from database
		_, err := db.Exec(`DELETE FROM users WHERE id = ?`, 1)
		if err != nil {
			log.Fatal(err)
		}
	}
}
