package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"io"
)

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576 // 1 Mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != io.EOF {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != nil {
		return errors.New("body must only contain a single JSON object")
	}
	return nil
}

func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		_, err = w.Write(out)
		if err != nil {
			return err
		}
	}
	return nil
}