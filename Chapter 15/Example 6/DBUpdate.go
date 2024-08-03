// Updating existing data
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
	defer db.Close()

	UpdateStatement := `
		UPDATE test
		SET name = $1
		WHERE id = $2
	`
	updateResult, updateResultErr := db.Exec(UpdateStatement, "well,\nid = 4", "2")
	if updateResultErr != nil {
		panic(updateResultErr)
	}
	updatedRecords, updatedRecordsErr := updateResult.RowsAffected()
	if updatedRecordsErr != nil {
		panic(updatedRecordsErr)
	}
	fmt.Println("Number of records updated: ", updatedRecords)
}
