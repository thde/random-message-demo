package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("DBUSER"),
		Passwd:    os.Getenv("DBPASS"),
		Net:       "tcp",
		Addr:      "mysql-844ee58.3a76d95.mysql.nineapis.ch",
		DBName:    "app_prod",
		TLSConfig: "skip-verify", // skip verifying TLS Cert, it is selfsigned
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	defer db.Close() // Ensure database connection is closed even if there's an error

	// Execute the query to retrieve a row from DB
	var message string
	query := "SELECT message FROM random_messages ORDER BY RAND() LIMIT 1;"
	row := db.QueryRow(query)
	err = row.Scan(&message)
	if err != nil {
		log.Fatal(err)
	}

	// Print the retrieved message
	fmt.Println("Daily roasting:", message)
}
