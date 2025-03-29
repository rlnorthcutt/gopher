# Bootstrap (core/bootstrap.go)

The `Bootstrap()` function initializes all core services and systems for a GoPHER-based application. It wires together Echo, PocketBase, logging, tracing, caching, routing, jobs, permissions, and more.

Use it from `main.go` to spin up a ready-to-use server with minimal setup.

---

## ✅ What It Sets Up

| Component      | Description |
|----------------|-------------|
| Echo           | Web server, routing, middleware |
| PocketBase     | Auth, DB, storage, admin UI |
| Logger         | Structured logging via `zerolog` |
| Tracer         | OpenTelemetry-compatible tracer |
| Cache          | Ristretto cache with tag support |
| Jobs           | Scheduled tasks via `gocron` |
| Router         | Core route system and introspection |
| Metrics        | Prometheus/OpenTelemetry-compatible route |
| Permissions    | PocketBase rule sync manager |

---

## 🚀 Usage in `main.go`

```go
func main() {
  e, services, err := core.Bootstrap()
  if err != nil {
    log.Fatal(err)
  }

  core.InitAll(services)
  services.Jobs.Start()

  e.Logger.Fatal(e.Start(":8080"))
}
```

---

## 🧠 Internals

```go
func Bootstrap() (*echo.Echo, *CoreServices, error) {
  config := LoadConfig()
  logger := SetupLogger(config.AppName)
  tracer, _ := SetupTracer(config.AppName)

  e := echo.New()
  pb := pocketbase.New()
  cache := NewCache()
  jobs := NewJobScheduler(nil)

  services := New(e, pb, logger, tracer, config, cache, jobs)
  jobs.services = services

  e.Use(RequestIDMiddleware())
  e.Use(LoggingMiddleware(logger))
  InitRouter(e)
  AttachMetricsRoute(e)

  return e, services, nil
}
```

---

## 🧩 When to Use

- Every GoPHER starter or app should call `core.Bootstrap()` first
- You can swap pieces by customizing the return values
- Useful in CLI tools and jobs, too (not just HTTP servers)

---

## 🛣 Future Enhancements

- Env-based debug toggle
- Graceful shutdown helpers
- Service override support (e.g., mock DB or logger)

---

[← Back to README](../README.md)

