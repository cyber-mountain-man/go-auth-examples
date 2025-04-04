# Secure API Key Authentication (Go)

This project demonstrates a secure and flexible approach to API key authentication in Go using `.env` files to manage secrets. It builds on the basic version by separating sensitive data from source code, following best practices for environment-based configuration.

---

## 🔐 Features

- API keys are loaded securely from a `.env` file
- Supports multiple authorized API keys
- Routes are protected via reusable middleware
- JSON responses for all endpoints
- Real-world structure ideal for APIs and microservices

---

## 📁 Project Structure

```
secure/
├── main.go       # Main API server
├── .env          # Stores valid API keys (excluded from version control)
├── .gitignore    # Prevents .env from being pushed
├── go.mod        # Go module definition
```

---

## 🛡️ About `.gitignore` and `.env`

To follow security best practices, the `.env` file is listed in the `.gitignore` file so it will never be pushed to GitHub:

```
.env
```

This prevents accidental exposure of sensitive API keys.  
**Never commit secret values like API keys or credentials to your repository.**

If you're using version control with collaborators, share `.env` values securely (e.g., password manager or secure notes).

---

## 🧪 Getting Started

### Requirements

- Go 1.18+
- `github.com/joho/godotenv` (included via `go mod`)
- curl or Postman for testing

---

### 🔧 Setup

1. Clone the project and navigate into the `secure/` directory

2. Create a `.env` file:

```
VALID_API_KEYS=key123,key456,super-secret-key
```

3. Run the server:

```bash
go run main.go
```

---

## 📮 API Endpoints

### `GET /public`  
> Public route, no authentication required.

**Example:**

```bash
curl http://localhost:8080/public
```

**Response:**

```json
{
  "message": "Welcome to the secure public API endpoint!"
}
```

---

### `GET /data` (Protected)  
> Requires a valid API key in the `X-API-Key` header.

**Example:**

```bash
curl http://localhost:8080/data -H "X-API-Key: key123"
```

**Response:**

```json
{
  "message": "You have access to secure protected data!"
}
```

**Invalid Key Example:**

```bash
curl http://localhost:8080/data -H "X-API-Key: invalidkey"
```

**Response:**

```text
Unauthorized: missing or invalid API key
```

---

## 🔐 Security Notes

- API keys should never be committed to version control.
- `.env` is listed in `.gitignore` to prevent accidental leaks.
- In production, use tools like Docker Secrets or cloud-based secret managers.

---

## 🧠 Learning Goals

- Understand how to apply API key authentication
- Learn to separate config/secrets from code
- Use Go middleware to secure routes

---

## 📄 License

MIT License

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
```