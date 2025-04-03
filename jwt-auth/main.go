package main

import (
	"encoding/json" // For encoding/decoding JSON
	"fmt"           // For printing messages to console
	"net/http" // For building HTTP server and handling requests
	"time"     // For handling expiration time on tokens
	"github.com/golang-jwt/jwt/v5" // For working with JWT tokens
)

// Secret key used to sign JWTs (keep this private in production)
// In real apps, store it in an environment variable or config file.
var jwtKey = []byte("my_secret_key")

// Struct to hold login credentials (used to decode JSON from client)
type Credentials struct {
	Username string `json:"username"` // Extracts from JSON field "username"
	Password string `json:"password"` // Extracts from JSON field "password"
}

// Custom claims that will be embedded into the JWT
// This includes a username and standard JWT claims like expiration time.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims // Provides fields like ExpiresAt
}

func main() {
	// Route for logging in (generates token)
	http.HandleFunc("/login", loginHandler)

	// Protected route that requires JWT in the Authorization header
	http.HandleFunc("/dashboard", authMiddleware(dashboardHandler))

	// Start the server
	fmt.Println("JWT API server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// loginHandler handles POST /login requests to authenticate users
// If successful, returns a signed JWT token in the response body
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON body into creds struct
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)

	// Basic credential check — replace with DB lookup in real apps
	if err != nil || creds.Username != "admin" || creds.Password != "1234" {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Set token expiration to 15 minutes from now
	expirationTime := time.Now().Add(15 * time.Minute)

	// Create the JWT claims, which includes the username and expiry
	claims := &Claims{
		Username: creds.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Required for expiration check
		},
	}

	// Create a new JWT token using the HS256 signing method and our claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using our secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	// Return the token in JSON format
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

// authMiddleware protects routes by verifying the JWT token
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Read the "Authorization" header (e.g., "Bearer <token>")
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		// Strip "Bearer " from the header value to get the raw token
		tokenString := authHeader[len("Bearer "):]

		// Create a Claims object to hold the decoded claims
		claims := &Claims{}

		// Parse and validate the JWT token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // Use our secret key to validate the token
		})

		// If token is invalid or expired, reject the request
		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// If valid, allow the next handler to process the request
		next.ServeHTTP(w, r)
	}
}

// dashboardHandler is a protected endpoint that requires a valid JWT
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	// Respond with a success image in JSON format
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Welcome to your dashboard!",
	})
}