package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"vue-api/internal/data"
	"vue-api/internal/driver"
)

type Config struct {
	Port int
}

type Application struct {
	Config      Config
	InfoLog     *log.Logger
	ErrorLog    *log.Logger
	models      data.Models
	environment string
}

func main() {
	var cfg Config
	cfg.Port = 8081

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	dsn := os.Getenv("DSN")
	environment := os.Getenv("ENV")
	database, err := driver.ConnectPostgres(dsn)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	defer database.SQL.Close()

	app := &Application{
		Config:      cfg,
		InfoLog:     infoLog,
		ErrorLog:    errorLog,
		models:      data.New(database.SQL),
		environment: environment,
	}

	app.serve()

}

func (app *Application) serve() error {
	app.InfoLog.Println("API listening on port", app.Config.Port)

	serv := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.Config.Port),
		Handler: app.routes(),
	}

	return serv.ListenAndServe()
}
