package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v4"
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

	flag.Parse()

	ctx := context.Background()

	dsn := flag.String("dsn", "postgres://postgres:root@localhost:5432/snippetbox", "PostgreSQL data source name")

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn, ctx)
	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close(ctx)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &postgresql.SnippetModel{DB: db, Ctx: ctx},
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

func openDB(dsn string, ctx context.Context) (*pgx.Conn, error) {
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, err
	}
	//testing connection
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}
	//returning connection pool
	return conn, nil
}
