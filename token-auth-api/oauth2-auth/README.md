# Google OAuth2 Authentication (Go)

This project demonstrates a secure and minimal OAuth2 login flow using **Google Sign-In** with the Go programming language. It handles the full authentication process, including redirect, token exchange, session management, and user data retrieval — all while keeping sensitive credentials safely in a `.env` file.

---

## 🔐 Features

- Google OAuth2 login integration
- Secure session storage using encrypted cookies
- Loads credentials from `.env` (not hardcoded)
- Displays authenticated user’s name and email
- Clean and beginner-friendly Go code with comments

---

## 🛠 Tech Stack

- Go (net/http)
- Google OAuth2 API
- `gorilla/sessions` – for session management
- `godotenv` – for loading environment variables
- `golang.org/x/oauth2` – official OAuth2 implementation

---

## 📁 Project Structure

```
oauth2-auth/
├── main.go          # Core server logic with Google login and session handling
├── .env             # Stores your client ID/secret (never committed)
├── .gitignore       # Ensures .env and local artifacts are excluded from Git
├── go.mod           # Module file for dependencies
```

---

## 🔧 Setup Instructions

### 1. Create OAuth Credentials

- Go to [Google Cloud Console](https://console.cloud.google.com/)
- Create a new project (or use an existing one)
- Enable **Google People API**
- Create an **OAuth2 Client ID**:
  - Application Type: **Web application**
  - Redirect URI:  
    `http://localhost:8080/auth/google/callback`

---

### 2. Create `.env` File

Inside the `oauth2-auth/` folder:

```env
GOOGLE_CLIENT_ID=your-client-id-here
GOOGLE_CLIENT_SECRET=your-client-secret-here
GOOGLE_REDIRECT_URL=http://localhost:8080/auth/google/callback
```

⚠️ **Never commit your `.env` to version control**

---

### 3. Install Dependencies

```bash
go mod tidy
```

---

### 4. Run the App

```bash
go run main.go
```

Visit [http://localhost:8080](http://localhost:8080) in your browser and click **"Login with Google"**.

---

## 🧪 Example Response

After logging in:

```
✅ Logged in as: John Doe (johndoe@example.com)
[Logout Link]
```

---

## 🔐 Security Notes

- Secrets are stored in `.env` and excluded from Git
- `gorilla/sessions` uses encrypted cookie storage
- Always use HTTPS in production
- Consider validating the `state` parameter for CSRF protection

---

## 🧠 Learning Goals

- Understand the OAuth2 authorization code flow
- Learn how to manage sessions securely in Go
- See how to safely inject secrets using environment variables

---

## 📄 License

MIT License

---

**Author:** Guillermo Morrison  
_M.S. Information Systems & Cybersecurity Candidate | Backend Developer & Security Enthusiast_
```