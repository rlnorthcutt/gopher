# Role-Based Access Middleware

The `RequireRoles()` middleware provides route-level authorization based on user roles stored in PocketBase user records.

This allows you to restrict access to certain endpoints by requiring users to have one or more specific roles.

---

## ✅ Features

- Role-based authorization for authenticated users
- Simple declaration of required roles per route
- Returns 403 Forbidden if user lacks permission

---

## 📦 Usage

### Single Role
```go
admin := e.Group("/admin")
admin.Use(core.AuthMiddleware())
admin.Use(core.RequireRoles("admin"))
```

### Multiple Roles
```go
team := e.Group("/dashboard")
team.Use(core.AuthMiddleware())
team.Use(core.RequireRoles("admin", "editor"))
```

---

## 🔍 How It Works

1. Requires `AuthMiddleware()` to run first
2. Retrieves user from request context
3. Looks for `role` field on user record
4. Compares to required roles
5. If match: continue
6. If no match: return `403 Forbidden`

---

## 🧠 Assumptions

- Your PB `users` collection includes a field called `role` (string)
- Roles are stored as lowercase identifiers (e.g., `admin`, `editor`, `viewer`)

---

## 🔒 Example User Record

```json
{
  "id": "abc123",
  "email": "user@example.com",
  "role": "editor"
}
```

---

## 🛣 Future Enhancements

- Support for role arrays (multi-role users)
- Permission-based middleware (`RequirePermission("posts.edit")`)
- Admin override flag or bypass

---

[← Back to README](../../README.md)