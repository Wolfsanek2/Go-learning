// holding user data in a table
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
		fmt.Println("The connection to the DB was initialized")
	}
	defer db.Close()
	err = createTable(db)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table was created")
	}

	addUserStatement, err := db.Prepare("INSERT INTO Users VALUES ($1, $2, $3)")
	if err != nil {
		panic(err)
	}
	err = addUser(addUserStatement, 1, "Daniel", "daniel@packt.com")
	if err != nil {
		panic(err)
	}
	err = addUser(addUserStatement, 2, "Florian", "florian@packt.com")
	if err != nil {
		panic(err)
	}
	addUserStatement.Close()

	updateEmailString := `
		UPDATE Users
		SET Email = $1
		WHERE ID = $2
	`
	updateEmailStatement, err := db.Prepare(updateEmailString)
	if err != nil {
		panic(err)
	}
	_, err = updateEmailStatement.Exec("user@packt.com", 1)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Email was changed")
	}
	updateEmailStatement.Close()

	removeUserString := `
		DELETE FROM Users
		WHERE ID = $1
	`
	removeUserStatement, err := db.Prepare(removeUserString)
	if err != nil {
		panic(err)
	}
	_, err = removeUserStatement.Exec(2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The second user was removed")
	}
	removeUserStatement.Close()
}

func createTable(db *sql.DB) error {
	createStatement := `
		CREATE TABLE public.Users (
			ID integer,
			Name character varying COLLATE pg_catalog."default",
			Email character varying
		)
		WITH (
			OIDS = FALSE
		)
	`
	_, err := db.Exec(createStatement)
	if err != nil {
		return err
	}
	return nil
}

func addUser(stmt *sql.Stmt, ID int, Name, Email string) error {
	_, err := stmt.Exec(ID, Name, Email)
	if err != nil {
		return err
	}
	fmt.Printf("The user with name: %s and email: %s was successfuly added\n", Name, Email)
	return nil
}
