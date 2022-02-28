package main

import (
	"flag"
	"fmt"
	"github.com/mukeshpilaniya/auth/internal/drivers"
	"github.com/mukeshpilaniya/auth/internal/models"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version = "1.0.0"
)

type APIConfig struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config      APIConfig
	infoLogger  *log.Logger
	errorLogger *log.Logger
	version     string
	DB          models.DBModel
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLogger.Println(fmt.Sprintf("Starting Backend Server in %s mode on port %d", app.config.env, app.config.port))

	return srv.ListenAndServe()
}

func main() {
	var cfg APIConfig
	flag.IntVar(&cfg.port, "port", 8081, "Server port to listen on")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development | production | maintenance}")
	flag.StringVar(&cfg.db.dsn, "dsn", "postgres://pilaniya:pilaniya@localhost:5432/auth?sslmode=disable", "DSN")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	conn, err := drivers.DBConnection(cfg.db.dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer conn.Close()

	app := &application{
		config:      cfg,
		infoLogger:  infoLog,
		errorLogger: errorLog,
		version:     version,
		DB: models.DBModel{
			DB: conn,
		},
	}

	err = app.serve()

	if err != nil {
		log.Fatalf(fmt.Sprintln(err))
	}
}


