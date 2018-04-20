// use "sqlite3 sqlite.db" to access from cmd
package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

func main() 
	database, _ := sql.Open("sqlite3", "./sqlite.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Jim", "Beam")
	rows, _ := database.Query("Select id, firstname, lastname FROM people")
	var id int
	var firstname, lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + ": " + lastname)
	}
}
