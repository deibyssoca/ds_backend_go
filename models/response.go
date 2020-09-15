package models

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Response data type
type Response struct {
	Status      int         `json:status`
	Data        interface{} `json:data`
	Message     string      `json:message`
	contentType string
	writer      http.ResponseWriter
}

// CreateDefaultResponse return default
func CreateDefaultResponse(w http.ResponseWriter) Response {
	return Response{Status: http.StatusOK, writer: w, contentType: "application/json"}
}

// NotFound Resource not dound
func (r *Response) NotFound() {
	r.Status = http.StatusNotFound
	r.Message = "Resource Not Found."
}

// SendNotFound method
func SendNotFound(w http.ResponseWriter) {
	response := CreateDefaultResponse(w)
	response.NotFound()
	response.Send()
}

// Send method
func (r *Response) Send() {
	r.writer.Header().Set("Content-Type", r.contentType)
	r.writer.WriteHeader(r.Status)

	output, _ := json.Marshal(&r)

	fmt.Fprintf(r.writer, string(output))
}

// SendData method
func SendData(w http.ResponseWriter, data interface{}) {
	response := CreateDefaultResponse(w)
	response.Data = data
	response.Send()
}
