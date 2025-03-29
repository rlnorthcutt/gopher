# Observability (core/observability.go)

The observability layer provides logging, tracing, and metrics support through structured logging (via `zerolog`) and distributed tracing (via OpenTelemetry).

This layer ensures all core services and requests are observable by default.

---

## ✅ Features

- Structured logging with `zerolog`
- OpenTelemetry tracing support
- Request ID middleware for correlation
- Logs enriched with context: trace ID, request ID, duration

---

## 📦 Logger Setup

```go
func SetupLogger(appName string) zerolog.Logger
```

This returns a structured logger instance:

```go
logger := core.SetupLogger("myapp")
logger.Info().Msg("Booting up...")
```

---

## 🎯 Tracing

```go
func SetupTracer(appName string) (trace.Tracer, error)
```

This creates an OpenTelemetry-compatible tracer:

```go
tracer, err := core.SetupTracer("myapp")
spanCtx, span := tracer.Start(ctx, "RenderPage")
defer span.End()
```

---

## 🧠 Middleware: RequestID + Logging

### RequestID Middleware
Adds a unique ID to each request for traceability.

```go
e.Use(core.RequestIDMiddleware())
```

### Logging Middleware
Logs HTTP method, path, duration, and request ID.

```go
e.Use(core.LoggingMiddleware(logger))
```

Output:
```json
{
  "method": "GET",
  "path": "/about",
  "duration_ms": 48,
  "request_id": "a12d-fg..."
}
```

---

## 📊 Metrics

A Prometheus/OpenTelemetry-compatible route is registered:

```go
core.AttachMetricsRoute(e)
```

Exposes metrics at `/metrics` by default.

---

## 🛣 Future Enhancements

- Span propagation between services
- Logging to file/remote log collector
- Request sampling / slow log filters

---

[← Back to README](../README.md)