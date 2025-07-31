package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)


var Db *sql.DB;

func InitDB() {
	var err error;

	// modify the DB credentials to use .env 
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")); 
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
