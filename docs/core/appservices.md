# AppServices (core/appservices.go)

`AppServices` is the central container that holds all shared systems in the GoPHER framework. It is passed around to modules, jobs, and rendering logic to access things like logging, caching, config, and the PocketBase instance.

---

## ✅ Included Services

| Field        | Description |
|--------------|-------------|
| `Echo`       | The Echo server instance |
| `PB`         | PocketBase app instance |
| `Logger`     | Structured logger (zerolog) |
| `Tracer`     | OpenTelemetry tracer |
| `Cache`      | Ristretto in-memory cache |
| `Config`     | App configuration struct |
| `Jobs`       | Job scheduler wrapper |
| `Data`       | Data access service for PB |
| `Permissions`| PB collection rule manager |

---

## 🧱 Creating AppServices

This is typically created in `bootstrap.go`:

```go
services := core.New(echo, pb, logger, tracer, config, cache, jobs)
```

It automatically wires up:
- Entity system (`RegisterEntityRuntime`)
- Permissions manager
- Core service interfaces

---

## 🛠 Access Pattern

Pass `services` to anything that needs it:

```go
func (m *BlogModule) Init(services *core.AppServices) {
    services.Logger.Info().Msg("Blog module loaded")
}
```

Inside handlers:

```go
core.Services().Logger.Info().Msg("View page")
```

> Optional: use `core.Services()` for global access in simple apps.

---

## 🔁 Lifecycle

- Created once during `Bootstrap()`
- Passed into `InitAll()` and every module lifecycle method
- Also injected into jobs and templates

---

## 🛣 Future Enhancements

- Service override injection for testing/mocking
- Context-scoped service access
- CLI and job contexts separate from HTTP

---

[← Back to README](../README.md)

