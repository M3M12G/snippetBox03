package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/http"
	"os"
	"se03.com/pkg/models/postgresql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *postgresql.SnippetModel
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "user=postgres password=root host=localhost port=5432 dbname=snippetbox", "PostgreSQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgresql.SnippetModel{DB: db},
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*pgxpool.Pool, error) {
	conn, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return nil, err
	}
	//returning connection pool
	return conn, nil
}
