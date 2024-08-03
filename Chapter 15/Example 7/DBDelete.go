// Deleting data
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

	DeleteStatement := `
		DELETE FROM test
		WHERE id = $1
	`
	deleteResult, deleteResultErr := db.Exec(DeleteStatement, 2)
	if deleteResultErr != nil {
		panic(deleteResultErr)
	}

	deletedRecords, deletedRecordsErr := deleteResult.RowsAffected()
	if deletedRecordsErr != nil {
		panic(deletedRecords)
	}
	fmt.Println("Number of records deleted:", deletedRecords)
}
