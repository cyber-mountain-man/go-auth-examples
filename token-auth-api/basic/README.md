# Basic API Key Authentication (Go)

This project demonstrates a simple API key authentication mechanism using Go’s built-in `net/http` package. It is part of the [`go-auth-examples`](https://github.com/cyber-mountain-man/go-auth-examples) repository, which showcases various authentication techniques.

This example is ideal for beginners looking to understand how middleware works and how to protect API endpoints using custom headers.

---

## 🔐 What It Does

- Validates requests using a hardcoded list of API keys
- Protects selected endpoints via middleware
- Returns JSON responses for both public and protected routes
- Demonstrates a lightweight, beginner-friendly auth pattern

---

## 📁 Project Structure

```
basic/
├── main.go       # Main API server with middleware and route definitions
└── README.md     # Project documentation
```

---

## 🚀 Getting Started

### Requirements

- Go 1.18 or later
- A tool to test HTTP requests (e.g., curl, Postman)

### Run the Server

```bash
go run main.go
```

Server will start at:  
`http://localhost:8080`

---

## 📮 API Endpoints

### `GET /public`

- 🔓 Open to anyone — no authentication required

**Example:**

```bash
curl http://localhost:8080/public
```

**Response:**

```json
{
  "message": "Welcome to the public API endpoint!"
}
```

---

### `GET /data` (Protected)

- 🔐 Requires a valid API key in the `X-API-Key` header

**Valid API Keys (for demo purposes):**
- `12345`
- `abcdef`

**Example:**

```bash
curl http://localhost:8080/data -H "X-API-Key: 12345"
```

**Response:**

```json
{
  "message": "You have access to protected data!"
}
```

---

## 🧠 Key Concepts

| Concept              | Description |
|----------------------|-------------|
| **Middleware**        | Used to wrap protected routes and enforce authentication logic |
| **Custom Header**     | API key is passed in `X-API-Key` header |
| **Hardcoded Keys**    | Demonstrates logic without needing a database |
| **JSON API Responses**| Uses `encoding/json` to return structured data |

---

## 🛡️ Security Notes

> ⚠️ This is a learning example and **not production-ready** as-is.

- API keys are hardcoded in the source code — real systems should use:
  - Environment variables or a config file
  - Secure key storage and key rotation
- Rate limiting and logging are not implemented
- No HTTPS or encryption applied (important in production)

---

## 📄 License

This project is licensed under the MIT License.

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
```

---
