package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"

  )

func Connect() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=test_db_camp port=5432 sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	// connect to database using func `Connect`
	db, err := Connect()
	if err != nil {
	  panic(err)
	}
	// create table employee
	_, err = db.Exec(`CREATE TABLE employee (
		id INT,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address VARCHAR(255),
		salary INT
	  )`)
  
	if err != nil {
	  panic(err)
	}
  
	fmt.Println("Table employee created")
}


