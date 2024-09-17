package main

import (
	"net/http"
	"vue-api/internal/data"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (app *Application) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	mux.Get("/users/login", app.Login)
	mux.Post("/users/login", app.Login)

	mux.Get("/users/all", func(w http.ResponseWriter, r *http.Request) {
		var users data.User

		all, err := users.GetAll()
		if err != nil {
			app.ErrorLog.Println(err)
			return
		}

		app.writeJson(w, http.StatusOK, all)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "you@there.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}
		app.InfoLog.Println("Adding user")

		id, err := app.models.User.Insert(u)
		if err != nil {
			app.ErrorLog.Println(err)
			app.errorJSON(w, err, http.StatusForbidden)
			return
		}

		app.InfoLog.Println("Got back of", id)
		newUser, err := app.models.User.GetByID(id)
		if err != nil {
			app.ErrorLog.Println("cant get user", err)

		}
		app.writeJson(w, http.StatusOK, newUser)
	})

	return mux
}
