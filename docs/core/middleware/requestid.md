# Request ID Middleware

The `RequestIDMiddleware()` ensures that every HTTP request handled by Echo is assigned a unique identifier. This ID is used for logging, tracing, and debugging purposes.

---

## ✅ Features

- Generates a UUIDv4 for every request
- Adds the ID to request context and headers
- Useful for tracing logs across systems

---

## 🧱 Usage

Register it in your Echo instance:

```go
e.Use(core.RequestIDMiddleware())
```

---

## 📦 Behavior

- Adds a header `X-Request-ID` to the response
- Stores the ID in `c.Request().Context()`
- Used by `LoggingMiddleware()` and tracing spans

---

## 💡 Best Practices

- Use in all environments, including dev
- Include `request_id` in logs
- Pass it to external services for traceability

---

## 🛣 Future Enhancements

- Support for incoming `X-Request-ID` (to propagate from upstream)
- Store in Echo context helper for easy access

---

[← Back to README](../../README.md)