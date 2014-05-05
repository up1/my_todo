package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user1:user1password@/todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	//Insert data
	statementInsert, err := db.Prepare("INSERT INTO task VALUES( ?, ? )")
	if err != nil {
		panic(err.Error())
	}
	defer statementInsert.Close()

	result, err := statementInsert.Exec(0, "task 1")
	if err != nil {
		panic(err.Error())
	}

	lastID, err := result.LastInsertId()
	fmt.Printf("Last ID=%d\n", lastID)

	//Query data
	statementQuery, err := db.Prepare("SELECT task_id, title FROM task")
	if err != nil {
		panic(err.Error())
	}
	defer statementQuery.Close()

	rows, err := statementQuery.Query()
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		rows.Scan(&id, &title)
		fmt.Printf("ID=%d, Title=%s\n", id, title)
	}

}
