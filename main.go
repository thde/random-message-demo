package main

import (
	"cmp"
	"database/sql"
	_ "embed"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var (
	//go:embed "index.html"
	index string
	indexTmpl = template.Must(template.New("index").Parse(index))
)

const query = "SELECT message FROM messages ORDER BY RAND() LIMIT 1;"

type response struct {
	Message string
	Error   error
}

func run(logger *log.Logger) error {
	cfg := mysql.Config{
		User:      cmp.Or(os.Getenv("DB_USER"), "random_message_prod"),
		Passwd:    os.Getenv("DB_PASS"),
		Net:       "tcp",
		Addr:      os.Getenv("DB_HOST"),
		DBName:    cmp.Or(os.Getenv("DB_NAME"), "random_message_prod"),
		TLSConfig: "skip-verify", // skip verifying TLS Cert, it is selfsigned
	}
	logger.Printf("db config user=%q db=%q addr=%q", cfg.User, cfg.DBName, cfg.Addr)

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	}
	logger.Printf("db ping successful")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var message string
		row := db.QueryRow(query)
		err = row.Scan(&message)
		if err != nil {
			logger.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		err := indexTmpl.Execute(w, response{
			Message: message,
			Error:   err,
		})
		if err != nil {
			logger.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	s := &http.Server{
		Addr:         ":" + cmp.Or(os.Getenv("PORT"), "8080"),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Printf("listening on addr=%q", s.Addr)
	return s.ListenAndServe()
}

func main() {
	logger := log.New(os.Stderr, "", log.Flags())
	err := run(logger)
	if err != nil {
		logger.Fatal(err)
	}
}
