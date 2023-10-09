package utils

import (
	"io"
	"net/http"
	"encoding/json"
	"log"
)

// ParseBody reads the JSON data from the request body and parses it into the provided interface x.
func ParseBody(r *http.Request, x interface{}) {
	// Read the request body into a byte slice using io.ReadAll()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		// If there is an error reading the body, log the error and return.
		log.Printf("Error reading request body: %v", err)
		return
	}

	// Unmarshal the JSON-encoded data in the body to the provided interface x.
	err = json.Unmarshal(body, x)
	if err != nil {
		// If there is an error unmarshaling JSON data, log the error and return.
		log.Printf("Error unmarshaling JSON data: %v", err)
		return
	}

	// Log a message indicating successful parsing of the request body.
	log.Println("Request body parsed successfully.")
	// No need to return anything because x is modified directly as it's passed by reference.
}