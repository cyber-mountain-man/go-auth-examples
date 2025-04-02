# Cookie-Based Session Authentication (Go)

This project demonstrates a simple, secure cookie-based session authentication system built with Go using the `net/http` package and the `gorilla/sessions` middleware.

It is part of the [`go-auth-examples`](https://github.com/cyber-mountain-man/go-auth-examples) repository, which showcases various authentication methods for Go-based web applications.

---

## 🔐 Features

- Secure session management using encrypted cookies
- Basic login and logout functionality
- Route protection using middleware
- HTML templates for login, dashboard, and public pages
- Clear and concise structure for educational purposes

---

## 📁 Folder Structure

```
cookie-session-auth/
├── main.go                  # Main application file
├── go.mod                   # Go module definition
└── templates/               # HTML templates
    ├── home.html
    ├── about.html
    ├── login.html
    ├── dashboard.html
    ├── profile.html
    └── welcome.html
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher
- A modern browser (for testing)

### Run the Project

1. Navigate to the project directory:
   ```bash
   cd cookie-session-auth
   ```

2. Run the server:
   ```bash
   go run main.go
   ```

3. Visit `http://localhost:8080` in your browser.

---

## 🧪 Test Login Credentials

- **Username:** `admin`  
- **Password:** `1234`

Successful login redirects you to the protected dashboard and profile pages.

---

## 🔧 Technologies Used

- [Go](https://golang.org/) – Standard HTTP server
- [Gorilla Sessions](https://github.com/gorilla/sessions) – Secure cookie session management
- HTML5 templates – Rendered via Go’s `html/template` package

---

## 📚 Educational Purpose

This project is intended to be a learning resource for developers who are new to Go or want to understand how cookie-based session authentication works in a practical web application.

For more examples, visit the main repository:  
👉 [`go-auth-examples`](https://github.com/cyber-mountain-man/go-auth-examples)

---

## 📄 License

This project is licensed under the MIT License.

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
