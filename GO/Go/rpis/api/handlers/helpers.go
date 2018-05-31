package handlers

import (
	"fmt"
	"net/http"
)

func Response(rw http.ResponseWriter, statusCode int, message string) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	fmt.Fprintf(rw, message)
	return
}
