package main

import (
	"net/http"
	"time"
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

		payload := JsonResponse{
			Error:   false,
			Message: "success",
			Data:    envelope{"users": all},
		}

		app.writeJson(w, http.StatusOK, payload)
	})

	mux.Get("/users/add", func(w http.ResponseWriter, r *http.Request) {
		var u = data.User{
			Email:     "you@there.com",
			FirstName: "You",
			LastName:  "There",
			Password:  "password",
		}
		app.InfoLog.Println("Adding user...")

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

	mux.Get("/test-generate-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute)

		if err != nil {
			app.ErrorLog.Println(err)
			return
		}
		token.Email = "admin@example.com"
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		payload := JsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJson(w, http.StatusOK, payload)
	})

	mux.Get("/test-save-token", func(w http.ResponseWriter, r *http.Request) {
		token, err := app.models.User.Token.GenerateToken(2, 60*time.Minute)

		if err != nil {
			app.ErrorLog.Println(err)
			return
		}

		user, err := app.models.User.GetByID(2)

		if err != nil {
			app.ErrorLog.Println(err)
			return
		}

		token.UserID = user.ID
		token.CreatedAt = time.Now()
		token.UpdatedAt = time.Now()

		err = token.Insert(*token, *user)

		if err != nil {
			app.ErrorLog.Println(err)
			return
		}

		payload := JsonResponse{
			Error:   false,
			Message: "success",
			Data:    token,
		}

		app.writeJson(w, http.StatusOK, payload)
	})

	mux.Get("/test-validate-token", func(w http.ResponseWriter, r *http.Request) {
		tokenToValidate := r.URL.Query().Get("token")
		valid, err := app.models.Token.ValidToken(tokenToValidate)
		if err != nil {
			app.errorJSON(w, err)
			return
		}

		var payload JsonResponse
		payload.Error = false
		payload.Data = valid

		app.writeJson(w, http.StatusOK, payload)

	})

	return mux
}
