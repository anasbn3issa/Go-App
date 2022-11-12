package main

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Error bool `json:"error"`
	Message string `json:"message"`
}

func (app *application) Login(w http.ResponseWriter, r *http.Request) {
	type credentials struct {
		Nickname string `json:"nickname"`
		Password string `json:"password"`}

		var creds credentials
		var payload jsonResponse

		err := json.NewDecoder(r.Body).Decode(&creds)

		if err != nil {
			//send back error message
			app.errorLog.Println("Error decoding JSON", err)
			payload.Error = true
			payload.Message = "Error decoding JSON"

			out, err := json.MarshalIndent(payload, "", "\t")
			if err != nil {
				app.errorLog.Println(err)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(out)
		}

		// TODO : authenticate 
		app.infoLog.Println("User", creds.Nickname, "pwd", creds.Password)

		// send back a response
		payload.Error = false
		payload.Message = "Login successful"

		out, err := json.MarshalIndent(payload, "", "\t")
		if err != nil {
			app.errorLog.Println(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(out)

}

