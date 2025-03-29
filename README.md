# GoPHER Framework

**GoPHER** is a modular, extensible Go framework for building low-code, server-rendered SaaS apps with PocketBase, Echo, HTMX, and Ristretto.

## ✨ Core Stack

- **G**o — the core language
- **O**penTelemetry — structured logging and tracing
- **P**ocketBase — database, file storage, auth, and admin
- **H**TMX — modern frontend without the SPA complexity
- **E**cho — fast and composable web server
- **R**istretto — high-performance in-memory caching

## 🧩 Key Features

- 🔌 Modular architecture (like Drupal modules)
- 🧠 Smart caching with tag-based invalidation
- 🔐 Integrated auth (email, OAuth2, OTP, MFA)
- 🔍 Built-in logging, tracing, metrics
- 🛠 CLI + starter templates (planned)
- 🧱 Entity modeling with PocketBase + Go
- 🌐 HTMX + Go templates for server-rendered interactivity

---

## 🛠 How to Use

### 1. Create a Starter App

```bash
git clone https://github.com/yourname/gopher-starter myapp
cd myapp
go run main.go
```

### 2. Enable Modules

```go
func init() {
    core.Enable(&BlogModule{})
    core.Enable(&PagesModule{})
}
```

### 3. Define Routes in Modules

```go
core.Route.Register("GET", "/blog", m.Index, "blog.index", "blog")
```

### 4. Use Core Services

```go
core.RenderPage(c, "blog/view", data)
services.Cache.Set("blog:/home", html, 5*time.Minute)
```

---

## 📁 Core Directory Structure

```bash
core/
├── appservices.go     # AppServices struct and global wiring
├── bootstrap.go        # Bootstraps the app and core systems
├── cache.go            # Ristretto cache wrapper + tagging
├── config.go           # Environment and config loading
├── data.go             # PocketBase data helpers
├── entity.go           # Generic Go model interface
├── file.go             # File upload and access layer
├── jobs.go             # Scheduled job system
├── middleware/         # Request ID, logging, auth, etc
├── module.go           # Module lifecycle manager
├── observability.go    # Logger + OpenTelemetry setup
├── pb.go               # PocketBase client wrapper
├── permissions.go      # Collection-level rule manager
├── render.go           # Go template and HTMX rendering
├── router.go           # Central route registration
```

---

## 📚 Documentation

- [core/appcontext.go](docs/core/appservices.md)
- [core/bootstrap.go](docs/core/bootstrap.md)
- [core/cache.go](docs/core/cache.md)
- [core/config.go](docs/core/config.md)
- [core/data.go](docs/core/data.md)
- [core/entity.go](docs/core/entity.md)
- [core/file.go](docs/core/file.md)
- [core/jobs.go](docs/core/jobs.md)
- [core/module.go](docs/core/module.md)
- [core/observability.go](docs/core/observability.md)
- [core/pb.go](docs/core/pb.md)
- [core/permissions.go](docs/core/permissions.md)
- [core/render.go](docs/core/render.md)
- [core/router.go](docs/core/router.md)

Want to generate the first docs page next? (e.g. `docs/cache.md` or `docs/render.md`?)

