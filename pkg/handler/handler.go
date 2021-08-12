package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

// A (simple) example of our application-wide configuration.
type Config struct{}

// The Handler struct that takes a configured Env and a function matching
// our useful signature.
type Handler struct {
	*Config
	H func(w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
// Custom response http server
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(w, r)
	if err != nil {
		switch e := err.(type) {
		case Error:
			// We can retrieve the status here and write out a specific
			// HTTP status code.
			log.Printf("HTTP %d - %s", e.Status(), e)

			// Compose json error
			data, _ := json.Marshal(Response{
				Errors: Errors{
					Code:    http.StatusText(e.Status()),
					Message: e.Error(),
				},
			})

			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(e.Status())
			w.Write(data)
		default:
			// Any error types we don't specifically look out for default
			// to serving a HTTP 500
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
