# Auth Middleware

The `AuthMiddleware()` validates PocketBase authentication tokens and ensures that only authenticated users can access protected routes.

It supports bearer tokens from users, admins, or external apps.

---

## ✅ Features

- Validates PocketBase JWT tokens
- Refreshes token and returns updated auth info
- Rejects unauthenticated requests (401)
- Injects user/admin into request context

---

## 🔐 Usage

Apply to any Echo route group:

```go
secured := e.Group("/api")
secured.Use(core.AuthMiddleware())
```

---

## 🧱 How It Works

1. Extracts bearer token from `Authorization` header
2. Sends `authRefresh` request to PocketBase
3. If valid:
   - Saves user or admin into request context
   - Continues to handler
4. If invalid:
   - Returns `401 Unauthorized`

---

## 👤 Accessing Auth Context

Use this helper in your handler:

```go
user := core.GetCurrentUser(c)
if user != nil {
    fmt.Println("User ID:", user.Id)
}
```

Supports:
- `GetCurrentUser(c)` – for regular users
- `GetCurrentAdmin(c)` – for admins
- `GetCurrentAuth(c)` – for either

---

## 📦 Notes

- This middleware does not invalidate tokens
- You can ignore the refreshed token if not needed
- Tokens are not stored server-side by PocketBase

---

## 🛣 Future Enhancements

- Role-based access control middleware
- Session and IP audit tracking
- Caching authRefresh results for performance

---

[← Back to README](../../README.md)