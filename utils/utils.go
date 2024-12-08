package utils

import (
	"encoding/json" // For encoding and decoding JSON data
	"fmt"           // For formatted I/O operations
	"net/http"      // For HTTP request and response handling

	"github.com/go-playground/validator/v10" // For data validation
)

var Validate = validator.New()

// ParseJSON decodes a JSON-encoded request body into the given payload.
// It reads from the HTTP request body and populates the provided `payload`.
// If the request body is missing or there is an issue decoding the JSON,
// it returns an error.
func ParseJSON(r *http.Request, payload any) error {
	// Step 1: Check if the request body is nil
	if r.Body == nil {
		return fmt.Errorf("request body is missing") // Return an error if the body is missing
	}

	// Step 2: Decode the JSON into the provided payload
	return json.NewDecoder(r.Body).Decode(payload)
}

// WriteJSON encodes a given value into JSON format and writes it to the HTTP response.
// It sets the appropriate `Content-Type` header and the response status code,
// and then writes the encoded JSON to the response writer.
func WriteJSON(w http.ResponseWriter, status int, v any) error {
	// Step 1: Set the Content-Type header to "application/json"
	w.Header().Add("Content-Type", "application/json")
	// Step 2: Write the HTTP status code to the response
	w.WriteHeader(status)
	// Step 3: Encode the provided value into JSON and write it to the response body
	return json.NewEncoder(w).Encode(v)
}

// WriteError is a helper function to send error messages as a JSON response.
// It calls `WriteJSON` to encode and send the error message in JSON format.
// The `status` is the HTTP status code, and the `err` is the error that should be returned in the response body.
func WriteError(w http.ResponseWriter, status int, err error) error {
	// Step 1: Create a map containing the error message in a key-value pair
	// Step 2: Call WriteJSON to send the error as a JSON response
	return WriteJSON(w, status, map[string]string{"error": err.Error()})
}