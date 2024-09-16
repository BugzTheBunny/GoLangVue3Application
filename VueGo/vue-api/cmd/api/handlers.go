package main

import (
	"encoding/json"
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

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		app.ErrorLog.Println("invalid Json")
		payload.Error = true
		payload.Message = "invalid Json"
		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.ErrorLog.Println(err)
		}

		w.Header().Set("content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(out)
		return
	}

	// Authenticate
	app.InfoLog.Println(creds.UserName, creds.Password)

	// Send back response

	payload.Error = false
	payload.Message = "Signed In"

	out, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		app.ErrorLog.Println(err)
	}

	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
