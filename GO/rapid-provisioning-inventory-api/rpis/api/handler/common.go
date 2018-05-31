package handler

import (
	"net/http"

	"fmt"

	"encoding/json"
)

// Error is the extended method of http.Error
func Error(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "json/application; charset=utf-8")
	w.WriteHeader(code)
	var e interface{}
	switch err.(type) {
	case error:
		e = err.(error).Error()
	case fmt.Stringer:
		e = err.(fmt.Stringer).String()
	case []error:
		errorList := err.([]error)
		errs := make([]string, 1)
		for _, er := range errorList {
			errs = append(errs, er.(error).Error())
		}
		e = err
	default:
		e = err
	}
	rerr := struct {
		Err interface{} `json:"error"`
	}{
		e,
	}
	b, err := json.Marshal(rerr)
	if err != nil {
		return
	}
	fmt.Fprintln(w, string(b))
}
