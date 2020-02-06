package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type human struct {
	id       int
	name     string
	lastname string
}

func initHuman() *human {
	return &human{
		id:       0,
		name:     "",
		lastname: "",
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./persons.db")
	if err != nil {
		log.Fatal(err)
	}

	// creat a database
	// table with first and last name
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, name TEXT, lastname TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	//executes the statements, returns the structure Stmt
	if _, err := statement.Exec(); err != nil {
		log.Fatal(err)
	}

	statement, err = db.Prepare("INSERT INTO people (name, lastname) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	// executes the statements, and gives to the query/statement the values
	if _, err := statement.Exec("Lee", "Silva"); err != nil {
		log.Fatal(err)
	}
	if _, err := statement.Exec("Afonso", "Oliveira"); err != nil {
		log.Fatal(err)
	}
	if _, err := statement.Exec("Miguel", "Silva"); err != nil {
		log.Fatal(err)
	}

	// Query executes all types of query
	// Query returs a point to a structure named Row presenting some functions
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal(err)
	}
	hum := initHuman()
	// Scan copies the columns in the current row into the values pointed
	// at by dest. The number of values in dest must be the same as the
	// number of columns in Rows.
	for rows.Next() {
		if err := rows.Scan(&hum.id, &hum.name, &hum.lastname); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d : %s %s\n", hum.id, hum.name, hum.lastname)
	}
}
