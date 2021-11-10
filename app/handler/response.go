package handler

import (
	"encoding/json"
	"net/http"
	"transport-manager/m/v1/app/model"
)

// SuccessArrRespond -> response formatter
func SuccessArrRespond(fields []*model.Transport, writer http.ResponseWriter) {
	// var fields["status"] := "success"
	_, err := json.Marshal(fields)
	type data struct {
		People     []*model.Transport `json:"data"`
		Statuscode int                `json:"status"`
		Message    string             `json:"msg"`
	}
	temp := &data{People: fields, Statuscode: 200, Message: "success"}
	if err != nil {
		ServerErrResponse(err.Error(), writer)
	}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(temp)
}

// ServerErrResponse -> server error formatter
func ServerErrResponse(error string, writer http.ResponseWriter) {
	type servererrdata struct {
		Statuscode int    `json:"status"`
		Message    string `json:"msg"`
	}
	temp := &servererrdata{Statuscode: 500, Message: error}

	//Send header, status code and output to writer
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(writer).Encode(temp)
}
