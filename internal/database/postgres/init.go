package postgres

import (
	"database/sql"
	"fmt"
	"os"

	//for connecting to db

	_ "github.com/lib/pq"
)

const (
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	user     = os.Getenv("DB_USER")
	password = os.Getenv("DB_PASSWORD")
	dbname   = os.Getenv("DB_DBNAME")
)

var (
	DbClient *sql.DB
)

func init() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	DbClient, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	// defer dbClient.Close()
	fmt.Println((DbClient))
	fmt.Println("Successfully connected!")
}
