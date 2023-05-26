package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"time"
	"database/sql"
  )

  type User struct {
	ID           uint
	Name         string
	Email        string
	Age          uint8
	Birthday     string
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
  }

  func main() {
    dsn := "host=localhost user=postgres password=Kmzway87a@ dbname=test_db_camp port=5432 sslmode=disable TimeZone=Asia/Jakarta"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    db.AutoMigrate(&User{})


	fmt.Println("Migration Success")

}