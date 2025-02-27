package router

import (
	"encoding/json"
	"net/http"
)

var MainRouter = http.NewServeMux()

func ErrorJson(writer http.ResponseWriter, message string, code int) {
	response := map[string]interface{}{
		"message": message,
		"status":  code,
	}

	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(writer).Encode(response)

	if err != nil {
		http.Error(writer, "Internal server error", http.StatusInternalServerError)
		return
	}
	return
}
