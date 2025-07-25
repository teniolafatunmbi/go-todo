package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"

)


var Db *sql.DB;

func InitDB() {
	var err error;

	// modify the DB credentials to use .env 
	connectionString := "postgres://postgres:postgres@database:5432/go_todo?sslmode=disable"; 
	Db, err = sql.Open("postgres", connectionString);

	if err != nil {
		log.Fatal(err);
		panic(err);
	}

	if err = Db.Ping(); err != nil {
		log.Fatal("Failed to ping DB: ", err);
		panic(err);
	}

	fmt.Println("Connected to database")

}
