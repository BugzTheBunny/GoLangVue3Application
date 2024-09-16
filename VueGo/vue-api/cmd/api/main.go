package main

import (
	"encoding/json"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var payload struct {
			Okay    bool   `json:"ok"`
			Message string `json:"message"`
		}

		payload.Okay = true
		payload.Message = "Hello, World"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.ErrorLog.Println(err)
		}

		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
	})

	app.InfoLog.Println("API listening on port", app.Config.Port)
	return http.ListenAndServe(fmt.Sprintf(":%d", app.Config.Port), nil)
}
