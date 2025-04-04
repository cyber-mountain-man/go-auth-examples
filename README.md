# Go Authentication Examples

This repository provides a collection of well-structured, beginner-friendly projects that demonstrate various authentication techniques using the Go programming language. Each project is self-contained and serves as a practical starting point for developers interested in learning how to implement secure authentication in Go-based web applications and APIs.

---

## 🔐 Authentication Methods Included

- ✅ **Cookie-Based Session Authentication** (using Gorilla Sessions)
- ⏳ **JWT (JSON Web Token) Authentication** – _In Progress_
- ✅ **OAuth2 Login with Google**
- ✅ **Token-Based API Authentication**
  - Basic key-checking
  - Secure `.env`-driven validation

---

## 📂 Project Overview

Each example lives in its own folder and is organized for clarity and ease of learning:

| Folder | Description |
|--------|-------------|
| [`cookie-session-auth/`](./cookie-session-auth) | Cookie-based login system using Gorilla Sessions |
| [`jwt-auth/`](./jwt-auth) | JWT-based authentication API _(coming soon)_ |
| [`token-auth-api/basic`](./token-auth-api/basic) | Minimal API key auth example with hardcoded keys |
| [`token-auth-api/secure`](./token-auth-api/secure) | Secure API key auth using `.env` and `godotenv` |
| [`token-auth-api/oauth2-auth`](./token-auth-api/oauth2-auth) | Google OAuth2 login flow with user sessions |

---

## 🎯 Goals

- Provide real-world examples of common authentication patterns in Go
- Help new developers build secure web applications and APIs
- Offer a clean reference point for comparing authentication techniques

---

## 🚀 Getting Started

Each folder contains its own `README.md` with:

- 📦 Installation instructions
- 🧪 API endpoints or login URLs
- 🛡 Security best practices
- 🧠 Learning goals

Clone the repo and dive into any project!

```bash
git clone https://github.com/cyber-mountain-man/go-auth-examples.git
cd go-auth-examples/<project-folder>
go run main.go
```

---

## 🤝 Contributions

Contributions are welcome! Whether you'd like to improve an example, fix a bug, or add a new authentication method (e.g., GitHub OAuth, passwordless login, TOTP MFA), feel free to fork the repository and submit a pull request.

---

## 📄 License

MIT License

---

**Maintained by:** Guillermo Morrison  
_Graduate student in Information Systems & Cybersecurity | Passionate about backend development and secure system design_
