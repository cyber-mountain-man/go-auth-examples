

### 📄 `jwt-auth/README.md`

```markdown
# JWT-Based Authentication API in Go

This project demonstrates a basic JSON Web Token (JWT) authentication flow built with Go. It is part of the [`go-auth-examples`](https://github.com/cyber-mountain-man/go-auth-examples) repository, which contains a collection of authentication examples for learning and reference.

---

## 🔐 Features

- Stateless authentication using signed JWTs
- Secure token generation with expiration
- Middleware-protected routes
- JSON request and response handling
- Clear, minimal structure for educational use

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

#### Example Request with `curl`

##### Linux/macOS:

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"1234"}'
```

##### Windows PowerShell:

```powershell
curl -X POST http://localhost:8080/login `
  -H "Content-Type: application/json" `
  -d "{\"username\":\"admin\",\"password\":\"1234\"}"
```

> ✅ Tip: You can also use [Postman](https://www.postman.com/) for easier testing without needing to escape quotes.

---

### `GET /dashboard` _(Protected)_

Requires a valid JWT in the `Authorization` header.

#### Header Format

```
Authorization: Bearer <your-jwt-token>
```

#### Example:

```bash
curl http://localhost:8080/dashboard \
  -H "Authorization: Bearer <your-jwt-token>"
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
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) – JWT creation & validation
- `net/http` – Built-in web server & routing

---

## 🎯 Purpose

This project is intended for developers who want to learn how to:

- Implement stateless authentication in Go
- Secure API endpoints using JWTs
- Understand basic auth middleware patterns

It’s a foundational example designed for educational and prototyping purposes.

---

## 🛡️ Security Notes

- 🔒 The JWT signing key is hardcoded — in production, use environment variables
- 🧑 Replace the hardcoded user with proper user authentication (e.g., database)
- 🕒 Tokens expire after 15 minutes — implement refresh tokens for longer sessions

---

## 📄 License

This project is licensed under the MIT License.

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
```