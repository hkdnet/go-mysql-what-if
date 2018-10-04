package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" // use mysql
)

const connectionString = "root:password@tcp(mysql)/playground?parseTime=true"

func main() {
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	delQuery := "delete from foo;"
	_, err = db.Exec(delQuery)
	if err != nil {
		log.Fatal(err)
	}

	insQuery := "insert into foo (id, foo) values (?, ?);"
	_, err = db.Exec(insQuery, 1, 1)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(insQuery, 1, 2)
	if err != nil {
		log.Println("expected")
		log.Println(err) // Error 1062: Duplicate entry '1' for key 'PRIMARY'
	}
	_, err = db.Exec(insQuery, 2, 1)
	if err != nil {
		log.Println("expected")
		log.Println(err) // Error 1062: Duplicate entry '1' for key 'foo'
	}
}
