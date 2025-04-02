// Declare the main package - this is where the Go program starts.
package main

// Import the neccessary packages
import (
	"html/template" // Used for rendering HTML templates
	"net/http"      // Provides HTTP server functionality
	"github.com/gorilla/sessions" // External package for managing sessions and cookies
)

// Create a new cookie-based session store with a secret key.
// THis key is used to sign cookies to prevent tampering.
var store = sessions.NewCookieStore([]byte("super-secret-key"))

func main() {
	// Set up Http route handlers - these handle requests to different URLs

	http.HandleFunc("/", homeHandler)			// Public home page
	http.HandleFunc("/about", aboutHandler)		// Public about page
	http.HandleFunc("/login", loginHandler)		// Login page
	http.HandleFunc("/logout", logoutHandler)	// Logout action
	http.HandleFunc("/welcome", welcomeHandler)	// Simple welcome page

	// Routers below are protected and requrie the user to be logged in
	http.HandleFunc("/dashboard", authMiddleware(dashboardHandler)) // Protected dashboard
	http.HandleFunc("/profile", authMiddleware(profileHandler))		// Protected user profile

	// Start the server on port 8080.  nil means use the default multiplexer.
	http.ListenAndServe(":8080", nil)
}

// ===========================
// Public page handlers
// ===========================

// Displayes the home page - anyone can access this
func homeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "home.html") // Render the home.html template
}

// Displays the about page - anyone can access this
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "about.html") // Render the about.html template
}

// ===========================
// Login & Logout
// ===========================

// Handles the login form
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// Show the login page if it's a GET request
	if r.Method == "GET" {
		render(w, "login.html")
		return
	}

	// Process login submission (POST request)
	username := r.FormValue("username") // Get the username from form
	password := r.FormValue("password") // Get the password from form

	// Check if credentials are correct (this is hardcoded - because of no real user DB yet)
	if username == "admin" && password == "1234" {
		//Get or create a session for this user
		session, _ := store.Get(r, "session")

		// Mark user as authenticated
		session.Values["authenticated"] = true
		session.Save(r, w) // Save the session (writes a cookie to browser)

		// Redirect to the protected dashboard page
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}

	// If login fails, show an error
	http.Error(w, "Invalid login", http.StatusUnauthorized)
}

// Handles logout action
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session
	session, _ :=store.Get(r, "session")

	// Mark user as logged out
	session.Values["authenticated"] = false
	session.Save(r, w)

	// Redirect to home page
	http.Redirect(w, r, "/", http.StatusFound)
}

// ===========================
// Protected Page Handlers
// ===========================


// Simple welcome page - could be protected or public
func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "welcome.html")
}

// Only accesible if logged in
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "dashboard.html")
}

// Only accesible if logged in
func profileHandler(w http.ResponseWriter, r *http.Request) {
	render(w, "profile.html")
}


// ==============================
// Middleware for Auth Protection
// ==============================


// This middleware function wraps other handlers and checks if the user is logged in
func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Retrieve the session
		session, _ := store.Get(r, "session")

		// Check if "authenticated" is set to true in the session
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			// If not authenticated, redirect to login page
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		// If authenticated, allow access to the next handler
		next.ServeHTTP(w, r)
	}
} 


// ===========================
// Template Rendering Helper
// ===========================


// Renders HTML templates located in the "templates" folder
func render(w http.ResponseWriter, filename string) {
	// Load the specified HTML file
	tmpl, err := template.ParseFiles("templates/" + filename)

	// If something goes wrong, show an error
	if err != nil {
		http.Error(w, "Page not found", http.StatusInternalServerError)
		return
	}

	// Render the template with no dynamic data (nil for now)
	tmpl.Execute(w, nil)
}