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

    // update data to table `employees`
    _, err = db.Exec(`UPDATE employees SET salary = 1000000 WHERE id =1 `)

    if err != nil {
        panic(err)
    }

    fmt.Println("Update data success")


    // delete data to table `employees`
    _, err = db.Exec(`DELETE FROM employees WHERE id = 5`)
    if err != nil {
        panic(err)
    }

    fmt.Println("Delete data success")
}