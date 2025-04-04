package main

import (
	"encoding/json"  // For encoding data into JSON responses
	"fmt"            // For printing logs to the terminal
	"net/http"       // Core HTTP package for building web servers
	"os"             // Accesses environment variables
	"strings"        // Helps manipulate strings (like splitting comma-separated values)

	"github.com/joho/godotenv" // Library to load .env file into environment variables
)

// validAPIKeys is a global map storing authorized API keys.
// Keys are loaded from an external .env file instead of being hardcoded.
var validAPIKeys = map[string]bool{}

func main() {
	// ===== Step 1: Load Environment Variables =====
	// Load environment variables from a file named ".env"
	// This allows you to separate secrets (like API keys) from your codebase.
	err := godotenv.Load()
	if err != nil {
		// If .env file isn't found, log a warning (it's not fatal if keys are set in OS env)
		fmt.Println("Warning: .env file not found. Using system environment variables.")
	}

	// ===== Step 2: Read and Parse the API Keys =====
	// Read the API key list from the environment variable VALID_API_KEYS
	// Example format: "key123,key456,key789"
	apiKeyCSV := os.Getenv("VALID_API_KEYS")

	// Split the comma-separated string and load each key into the validAPIKeys map
	for _, key := range strings.Split(apiKeyCSV, ",") {
		trimmedKey := strings.TrimSpace(key)
		if trimmedKey != "" {
			validAPIKeys[trimmedKey] = true
		}
	}

	// ===== Step 3: Register Routes =====
	// Public route: no authentication required
	http.HandleFunc("/public", publicHandler)

	// Protected route: only accessible with a valid API key
	// Wrapped with the apiKeyMiddleware to enforce authentication
	http.HandleFunc("/data", apiKeyMiddleware(protectedHandler))

	// Log to terminal that the server is running
	fmt.Println("Secure API Key Auth server running at http://localhost:8080")

	// Start the HTTP server on port 8080
	http.ListenAndServe(":8080", nil)
}

// ===== Middleware: API Key Validator =====
// apiKeyMiddleware is a wrapper around protected routes that checks
// for a valid "X-API-Key" header in the request.
func apiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the "X-API-Key" header sent by the client
		key := r.Header.Get("X-API-Key")

		// Look up the key in the authorized map
		if _, ok := validAPIKeys[key]; !ok {
			// If key not found, deny access
			http.Error(w, "Unauthorized: missing or invalid API key", http.StatusUnauthorized)
			return
		}

		// If key is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	}
}

// ===== Public API Endpoint =====
// This endpoint is open to anyone, no API key needed.
func publicHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple JSON message
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to the secure public API endpoint!",
	})
}

// ===== Protected API Endpoint =====
// This endpoint is only accessible with a valid API key.
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with protected data
	json.NewEncoder(w).Encode(map[string]string{
		"message": "You have access to secure protected data!",
	})
}
