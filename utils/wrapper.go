package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// WrapAPIError wrapper for error response
func WrapAPIError(w http.ResponseWriter, r *http.Request, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	result, err := json.Marshal(map[string]interface{}{
		"Code":          code,
		"Error Type":    http.StatusText(code),
		"Error Details": message,
	})
	if err == nil {
		log.Printf("Request %s %s", r.Method, message)
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("Can't wrap API error : %s", err))
	}
}

// WrapAPISuccess wrapper for success response
func WrapAPISuccess(w http.ResponseWriter, r *http.Request, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	result, err := json.Marshal(map[string]interface{}{
		"Code":   code,
		"Status": message,
	})
	if err == nil {
		log.Printf("Request %s %s", r.Method, message)
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("Can't wrap API success : %s", err))
	}
}

// WrapAPIData wrapper fro data response
func WrapAPIData(w http.ResponseWriter, r *http.Request, data interface{}, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	result, err := json.Marshal(map[string]interface{}{
		"Code":   code,
		"Status": message,
		"Data":   data,
	})

	//result, err := json.Marshal(data)

	if err == nil {
		log.Printf("Request %s %s", r.Method, message)
		w.Write(result)
	} else {
		log.Println(fmt.Sprintf("Can't wrap API data : %s", err))
	}
}
