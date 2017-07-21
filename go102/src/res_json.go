package main

import (
	"encoding/json"
	"net/http"
)

func ToJSON(writer http.ResponseWriter, v interface{}) {

	output, err := json.Marshal(v)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(output)
}
