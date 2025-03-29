# Logging Middleware

The `LoggingMiddleware()` adds structured request logs to your Echo server using `zerolog`. It logs useful information like method, path, duration, status code, request ID, and trace ID.

---

## ✅ Features

- Structured JSON logging via `zerolog`
- Includes method, path, status, latency
- Logs request ID and OpenTelemetry trace ID
- Helps debug and monitor traffic

---

## 🧱 Usage

Register it in your app:

```go
logger := core.SetupLogger("myapp")
e.Use(core.LoggingMiddleware(logger))
```

---

## 📦 Log Output Example

```json
{
  "method": "GET",
  "path": "/about",
  "status": 200,
  "duration_ms": 48,
  "request_id": "abc123",
  "trace_id": "xyz789"
}
```

---

## ⚙️ Behavior

- Captures start time before handler
- Calculates latency after handler
- Logs at `Info` level by default

---

## 🔒 Trace ID Support

If OpenTelemetry is configured, it pulls the `trace_id` from the request context:

```go
if spanCtx := trace.SpanContextFromContext(req.Context()); spanCtx.IsValid() {
  event = event.Str("trace_id", spanCtx.TraceID().String())
}
```

---

## 🛣 Future Enhancements

- Custom log levels for routes
- Optional body/response size logging
- Colorized dev mode output

---

[← Back to README](../../README.md)