package main

import (
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Email string `json:"Email"`
		Password string `json:"password"`}

		var creds credentials
		var payload jsonResponse

		err := app.readJSON(w, r, &creds)
		if err != nil {
			payload.Error = true
			payload.Message = err.Error()
			_ = app.writeJSON(w, http.StatusBadRequest, payload)
		}

		// TODO : authenticate 
		app.infoLog.Println("User", creds.Email, "pwd", creds.Password)

		// send back a response
		payload.Error = false
		payload.Message = "Login successful"

		err = app.writeJSON(w, http.StatusOK, payload)

		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
}

