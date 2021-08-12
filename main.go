package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/zeihanaulia/go-error-handling/pkg/handler"
)

func indexHandler(rw http.ResponseWriter, r *http.Request) error {
	// simulate error
	if r.URL.Query().Get("err") == "yes" {
		return handler.StatusError{Code: http.StatusInternalServerError, Err: errors.New("internal server error")}
	}

	var resp SuccessRespose
	resp.Method = r.Method
	resp.Status = "ok"

	handler.RenderJSON(resp, rw)
	return nil
}

type SuccessRespose struct {
	Method string `json:"method,omitempty"`
	Status string `json:"status,omitempty"`
}

func main() {
	// register handler
	http.Handle("/", handler.Handler{Config: &handler.Config{}, H: indexHandler})

	// run http server
	log.Println("server ready on http://localhost:3030")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
