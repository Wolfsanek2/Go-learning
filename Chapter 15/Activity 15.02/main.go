// finding the messages of specific users
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Message struct {
	UserID  int
	Message string
}

func main() {
	db, err := sql.Open("postgres", "user=postgres password=PGpassword!123 host=127.0.0.1 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connection to the DB was initialized")
	}
	defer db.Close()

	createStatement := `
		CREATE TABLE Messages (
			UserID integer,
			Message character varying COLLATE pg_catalog."default"
		)
		WITH (
			OIDS = FALSE
		)
	`
	_, err = db.Exec(createStatement)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table was created")
	}

	messages := []Message{
		{1, "message 1"},
		{2, "message 2"},
		{3, "message 3"},
	}
	addMessageString := `
		INSERT INTO Messages VALUES ($1, $2)
	`
	addMessageStatement, err := db.Prepare(addMessageString)
	if err != nil {
		panic(err)
	}
	for _, message := range messages {
		_, err = addMessageStatement.Exec(message.UserID, message.Message)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("UserID: %d; Message: %s was added\n", message.UserID, message.Message)
		}
	}
	addMessageStatement.Close()

	searchID := 2
	searchMessageString := `
		SELECT Message From Messages WHERE UserID = $1
	`
	searchMessageStatement, err := db.Prepare(searchMessageString)
	if err != nil {
		panic(err)
	}
	result, err := searchMessageStatement.Query(searchID)
	if err != nil {
		panic(err)
	}
	var (
		message string
	)
	number := 0
	for result.Next() {
		err = result.Scan(&message)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Was found user with id: %d; message: %s", searchID, message)
		number++
	}
	if number == 0 {
		fmt.Println("The query returned nothing")
	}
	searchMessageStatement.Close()
}
