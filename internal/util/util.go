package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

var (
	ErrInvalidJSONBody = errors.New("body must have a single Json value")
)

// Payload is a struct for sending custom message
type Payload struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

// ReadJSON reads json from request body into data. only support a single json value in the body
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return ErrInvalidJSONBody
	}
	return nil
}

// WriteJSON write arbitrary data as Json and set json content type
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return nil
}

// BadRequest sends a Json response with status http.StatusBadRequest, describing the error
func BadRequest(w http.ResponseWriter, r *http.Request, err error) error {
	var payload Payload
	payload.Error=true
	payload.Message=err.Error()

	out, err :=json.MarshalIndent(&payload,"","\t")
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return nil
}

// InvalidCredentials send a Json response with status http.StatusUnauthorized, describing the error
func InvalidCredentials (w http.ResponseWriter) error{
	var payload Payload
	payload.Error=true
	payload.Message="invalid authentication credentials"

	out, err :=json.MarshalIndent(&payload,"","\t")
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return nil
}