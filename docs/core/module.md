# Module System (core/module.go)

The module system provides a plug-and-play architecture for building isolated features (like Pages, Blog, etc.) that can hook into the core framework lifecycle.

Each module can register routes, jobs, permissions, and perform migrations.

---

## ✅ Features

- Auto-registration via `core.Enable(...)`
- Lifecycle hooks: Init, Migrate, Update
- Route + job registration support
- Optional disabling support (planned)

---

## 📦 Module Interface

```go
type Module interface {
    Key() string
    Name() string
    Description() string
    Init(*AppServices) error
    Migrate() error
    Update() error
    RegisterRoutes(e *echo.Echo)
    RegisterJobs(j *JobScheduler)
}
```

---

## ✅ Enabling a Module

```go
func init() {
    core.Enable(&PagesModule{})
    core.Enable(&BlogModule{})
}
```

All modules must be enabled before calling `core.InitAll(services)`.

---

## 🧠 Lifecycle Methods

| Method         | Description |
|----------------|-------------|
| `Init()`       | Called on boot with core services injected |
| `Migrate()`    | Called after init to apply schema or permissions |
| `Update()`     | Optional: called for long-term upgrades |
| `RegisterRoutes()` | Bind Echo routes |
| `RegisterJobs()`   | Add cron jobs |

---

## 📁 Typical Module Layout

```bash
modules/pages/
├── module.go
├── permissions.yaml
├── templates/pages/view.html
```

---

## 🧩 Example Module Implementation

```go
type PagesModule struct {}

func (m *PagesModule) Key() string { return "pages" }
func (m *PagesModule) Name() string { return "Pages" }
func (m *PagesModule) Description() string { return "Public CMS pages" }

func (m *PagesModule) Init(s *core.AppServices) error { return nil }
func (m *PagesModule) Migrate() error {
    return core.Services().Permissions.LoadFromYAML("modules/pages/permissions.yaml")
}
func (m *PagesModule) RegisterRoutes(e *echo.Echo) {
    core.Route.Register("GET", "/:slug", m.HandlePage, "pages.view", "pages")
}
```

---

## 🛣 Future Enhancements

- Enable/disable modules at runtime
- Admin UI for module listing and info
- Dependencies between modules
- Test runner integration

---

[← Back to README](../README.md)