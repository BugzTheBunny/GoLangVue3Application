package main

import (
	"net/http"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {
	type Credentials struct {
		UserName string `json:"email"`
		Password string `json:"password"`
	}

	var creds Credentials
	var payload JsonResponse

	err := app.readJson(w, r, &creds)
	if err != nil {
		app.ErrorLog.Println(err)
		payload.Error = true
		payload.Message = "invalid json supplied, or json missing"
		_ = app.writeJson(w, http.StatusBadRequest, payload)
	}

	// Authenticate
	app.InfoLog.Println(creds.UserName, creds.Password)

	// Send back response
	payload.Error = false
	payload.Message = "Signed In"

	err = app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		app.ErrorLog.Println(err)
	}
}
