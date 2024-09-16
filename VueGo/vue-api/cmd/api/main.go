package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Port int
}

type Application struct {
	Config   Config
	InfoLog  *log.Logger
	ErrorLog *log.Logger
}

func main() {
	var cfg Config
	cfg.Port = 8081

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &Application{
		Config:   cfg,
		InfoLog:  infoLog,
		ErrorLog: errorLog,
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
