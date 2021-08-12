package handler

import (
	"encoding/json"
	"net/http"
)

func RenderJSON(resp interface{}, rw http.ResponseWriter) {
	var response Response
	response.Data = resp
	data, _ := json.Marshal(response)
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.Write(data)
}
