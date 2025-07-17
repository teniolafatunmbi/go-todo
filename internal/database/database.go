package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

)


var db *sql.DB;

func InitDB() {
	var err error;

	// modify the DB credentials to use .env 
	connectionString := "postgres://postgres:postgres@localhost:5455/postgres?sslmode=disable"; 
	db, err = sql.Open("postgres", connectionString);

	if err != nil {
		log.Fatal(err);
		panic(err);
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping DB: ", err);
		panic(err);
	}

	fmt.Println("Connected to database")

}
