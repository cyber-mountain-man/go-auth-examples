package main

import (
	"context"
	"encoding/gob"  // Needed to store custom structs in sessions
	"encoding/json" // For decoding Google’s user info response
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions" // Cookie-based session management
	"github.com/joho/godotenv"    // Loads .env into environment variables
	"golang.org/x/oauth2"         // OAuth2 core logic
	"golang.org/x/oauth2/google"  // Google-specific OAuth2 endpoint
)

// UserInfo defines what user data we’ll store from Google’s user info API.
type UserInfo struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Store session data using encrypted cookies.
// You should replace the key below with a secure, random value in production.
var store = sessions.NewCookieStore([]byte("super-secret-session-key"))

// Global OAuth2 configuration variable (populated in init)
var googleOAuthConfig *oauth2.Config

func init() {
	// Allow storing UserInfo struct inside session cookies
	gob.Register(UserInfo{})

	// Load values from .env file into environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found. Falling back to system env vars.")
	}

	// Initialize the Google OAuth2 config using .env values
	googleOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"), // Google redirects here after login
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func main() {
	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/auth/google/callback", callbackHandler)
	http.HandleFunc("/logout", logoutHandler)

	fmt.Println("OAuth2 server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// homeHandler: Public homepage that shows either a login link or the logged-in user's name/email
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set response content type to HTML so the browser renders HTML tags
	w.Header().Set("Content-Type", "text/html")

	// Get the current session
	session, _ := store.Get(r, "session")

	// Try to retrieve user info from the session
	user, ok := session.Values["user"].(UserInfo)

	if ok {
		// User is logged in
		fmt.Fprintf(w, "Logged in as: %s (%s)<br><a href='/logout'>Logout</a>", user.Name, user.Email)
	} else {
		// User not logged in
		fmt.Fprint(w, "<a href='/login'>Login with Google</a>")
	}
}

// loginHandler: Starts the Google OAuth2 login flow
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Google’s OAuth2 consent page
	url := googleOAuthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// callbackHandler: Handles Google's redirect back to our app after login
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// Get the "code" query parameter from Google
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	// Exchange the code for an access token
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to create a new authenticated HTTP client
	client := googleOAuthConfig.Client(context.Background(), token)

	// Call Google's user info API to get the user's profile data
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Decode the JSON response into our UserInfo struct
	var userInfo UserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info", http.StatusInternalServerError)
		return
	}

	// Store the user info in the session
	session, _ := store.Get(r, "session")
	session.Values["user"] = userInfo
	session.Save(r, w)

	// Redirect user back to the homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// logoutHandler: Clears the session and logs the user out
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	// Invalidate the session by setting it to expire immediately
	session.Options.MaxAge = -1
	session.Save(r, w)

	// Redirect user to the homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
