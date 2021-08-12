package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func indexHandler(rw http.ResponseWriter, r *http.Request) {
	var resp Response

	if r.URL.Query().Get("err") == "yes" {
		resp.Errors.Code = http.StatusText(http.StatusInternalServerError)
		resp.Errors.Message = "something when wrong"
		data, _ := json.Marshal(resp)
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		rw.WriteHeader(http.StatusOK)
		rw.Write(data)
		return
	}

	resp.Data.Method = r.Method
	resp.Data.Status = "ok"
	data, _ := json.Marshal(resp)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	rw.Write(data)
}

type Response struct {
	Data   Data   `json:"data,omitempty"`
	Errors Errors `json:"errors,omitempty"`
}

type Data struct {
	Method string `json:"method"`
	Status string `json:"status"`
}

type Errors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func main() {
	// register handler
	http.HandleFunc("/", indexHandler)

	// run http server
	log.Println("server ready on http://localhost:3030")
	log.Fatal(http.ListenAndServe(":3030", nil))
}
