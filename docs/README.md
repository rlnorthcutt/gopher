# GoPHER Framework Documentation

Welcome to the internal documentation for the **GoPHER** framework. This system is designed for building modular, low-code SaaS applications using Go, PocketBase, Echo, and HTMX.

---

## 📁 Documentation Structure

All core service documentation lives in `docs/core/`, organized by filename and feature.

### 🔧 Core Systems

| Module         | Description |
|----------------|-------------|
| [appcontext](core/appservices.md)   | AppServices struct, shared app state |
| [bootstrap](core/bootstrap.md)     | Initializes the full framework runtime |
| [cache](core/cache.md)             | In-memory caching with tag support |
| [config](core/config.md)           | Environment-based app configuration |
| [data](core/data.md)               | PocketBase data access helper layer |
| [entity](core/entity.md)           | Struct-based modeling with hooks |
| [file](core/file.md)               | Upload and manage PB file storage |
| [jobs](core/jobs.md)               | Scheduled job system with logging |
| [module](core/module.md)           | Plugin/module lifecycle manager |
| [observability](core/observability.md) | Logger, tracing, and request metrics |
| [permissions](core/permissions.md) | PB rule-based access control management |
| [render](core/render.md)           | Go templates with HTMX and caching |
| [router](core/router.md)           | Centralized route registration and lookup |

### 🧩 Middleware

| Middleware                        | Description |
|----------------------------------|-------------|
| [requestid](core/middleware/requestid.md) | Adds unique request ID to each request |
| [logging](core/middleware/logging.md)     | Logs method, path, status, duration, trace ID |
| [auth](core/middleware/auth.md)           | Validates PB auth tokens for protected routes |
| [roles](core/middleware/roles.md)         | Restricts access by user role (e.g., admin) |

---

## 🧱 How to Use

Each page contains:
- Overview of the service/module
- Core APIs and usage
- Example code snippets
- Future enhancement ideas

Use this as a reference while developing new modules, customizing routes, or contributing to the framework itself.

---

## 🚧 Coming Soon

- Tutorials and recipes
- CLI reference
- Dev + production setup
- Admin module guide

---

[← Back to Project README](../README.md)

