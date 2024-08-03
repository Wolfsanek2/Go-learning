// Creating a new project
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=PGpassword!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the DB was successfully initialized!")
	}

	ErrConnectivity := db.Ping()
	if ErrConnectivity != nil {
		panic(ErrConnectivity)
	} else {
		fmt.Println("Good to go!")
	}

	defer db.Close()
}
