package main

import (
	"errors"
	"net/http"
	"time"
)

type JsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type envelope map[string]interface{}

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

	// lookup for user by email
	user, err := app.models.User.GetByEmail(creds.UserName)
	if err != nil {
		app.errorJSON(w, errors.New("invalid username or password"))
		return
	}

	// validate user password
	validPassword, err := user.PasswordMatches(creds.Password)
	if err != nil || !validPassword {
		app.errorJSON(w, errors.New("invalid password"))
		return
	}

	// generate a token if the user is valid
	token, err := app.models.Token.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	// save token to database
	err = app.models.Token.Insert(*token, *user)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	// Send back response
	payload = JsonResponse{
		Error:   false,
		Message: "logged in",
		Data:    envelope{"token": token, "user": user},
	}

	err = app.writeJson(w, http.StatusOK, payload)
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}
}

func (app *Application) Logout(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Token string `json:"token"`
	}

	err := app.readJson(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	err = app.models.Token.DeleteByToken(requestPayload.Token)
	if err != nil {
		app.errorJSON(w, errors.New("invalid json"))
		return
	}

	payload := JsonResponse{
		Error:   false,
		Message: "logged out",
	}

	_ = app.writeJson(w, http.StatusOK, payload)
}
