package main

import (
	"encoding/json" // Used to encode data to JSON for API responses
	"fmt"           // Used for printing to the console (logging)
	"net/http"      // Standard HTTP library for creating servers and handling routes
)

// ===== Simulated API Key Store =====

// In a real-world app, API keys would be stored securely in a database or
// loaded from an environment variable (not hardcoded in code).
// Here, we're using a simple map to simulate a list of valid API keys.
var validAPIKeys = map[string]bool{
	"12345":  true,
	"abcdef": true,
}

// ===== Middleware Function to Enforce API Key Authentication =====

// Middleware in Go is a function that wraps another HTTP handler and
// adds extra functionality — in this case, checking for an API key.
func apiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the "X-API-Key" header from the incoming request
		key := r.Header.Get("X-API-Key")

		// Check if the provided API key exists in the map of valid keys
		if _, ok := validAPIKeys[key]; !ok {
			// If the key is not valid or missing, respond with 401 Unauthorized
			http.Error(w, "Unauthorized: missing or invalid API key", http.StatusUnauthorized)
			return
		}

		// If valid, call the next handler in the chain
		next.ServeHTTP(w, r)
	}
}

// ===== Public Endpoint =====

// This is an open route that anyone can access — no API key required.
// It simulates a public-facing endpoint in an API.
func publicHandler(w http.ResponseWriter, r *http.Request) {
	// Encode and send a simple JSON response
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the public API endpoint!",
	})
}

// ===== 🔐 Protected Endpoint =====

// This route is protected by the API key middleware. <Only requests that 
// include a valid "X-API-Key" header will reach this handler.
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Return a simple JSON response indicating successful access
	json.NewEncoder(w).Encode(map[string]string{
		"message": "You have access to protected data!",
	})
}

// ===== Application Entry Point =====

func main() {
	// Register the public handler to the "/public" route (no API key required)
	http.HandleFunc("/public", publicHandler)

	// Register the protected handler at "/data", but wrap it with apiKeyMiddleware
	// so it requires a valid API key
	http.HandleFunc("/data", apiKeyMiddleware(protectedHandler))

	// Print a setup message to the console
	fmt.Println("API Key Auth Server running at http://localhost:8080")

	// Start the HTTP server server on port 8080
	http.ListenAndServe(":8080", nil)
}