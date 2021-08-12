package handler

type Response struct {
	Data   interface{} `json:"data,omitempty"`
	Errors Errors      `json:"errors,omitempty"`
}

type Data struct {
	Method string `json:"method"`
	Status string `json:"status"`
}

type Errors struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
