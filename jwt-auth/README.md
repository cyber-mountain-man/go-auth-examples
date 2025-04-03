# JWT-Based Authentication API in Go

This project demonstrates a basic JSON Web Token (JWT) authentication flow built with Go. It is part of the [`go-auth-examples`](https://github.com/cyber-mountain-man/go-auth-examples) repository, which contains a collection of authentication examples for learning and reference.

---

## 🔐 Features

- Stateless authentication using signed JWTs
- Secure token generation and expiration
- Middleware to protect routes
- JSON request and response handling
- Minimal, clear structure for educational purposes

---

## 📁 Project Structure

```
jwt-auth/
├── main.go       # Entry point for the API server
├── go.mod        # Go module file
└── README.md     # Project documentation
```

---

## 🚀 Getting Started

### Prerequisites

- Go 1.18 or higher
- A REST client like Postman, Insomnia, or `curl`

### Installation

1. Navigate to this project folder:

   ```bash
   cd jwt-auth
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the server:

   ```bash
   go run main.go
   ```

---

## 📮 API Endpoints

### `POST /login`

Authenticates the user and returns a signed JWT token.

#### Request (JSON)

```json
{
  "username": "admin",
  "password": "1234"
}
```

#### Response (JSON)

```json
{
  "token": "<your-jwt-token>"
}
```

---

### `GET /dashboard` _(Protected)_

Requires a valid JWT in the `Authorization` header.

#### Header

```
Authorization: Bearer <your-jwt-token>
```

#### Response

```json
{
  "message": "Welcome to your dashboard!"
}
```

---

## 🔧 Technologies Used

- [Go](https://golang.org/)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) for token creation and verification
- `net/http` standard library for routing and HTTP handling

---

## 🎯 Purpose

This project is designed for developers who want to learn how to:

- Implement stateless authentication in Go
- Work with JWT tokens
- Protect REST API routes
- Understand basic middleware usage

It serves as a foundational example for more advanced auth setups, such as refresh tokens and OAuth2.

---

## 🛡️ Security Notes

- The signing key (`jwtKey`) is hardcoded for demo purposes — store secrets securely in environment variables in production
- In real applications, replace the hardcoded user check with a proper user database

---

## 📄 License

This project is licensed under the MIT License.

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
```