# Router System (core/router.go)

The `Router` system provides a centralized API for registering, tracking, and managing routes in a modular and introspectable way.

It wraps the Echo router and adds metadata storage, dynamic lookup, and route introspection for docs, admin UIs, or CLI tools.

---

## ✅ Features

- Dynamic route registration
- Named route lookup (`Route.Get()`)
- Reverse routing (`Route.PathFor()`)
- Group support (by module)
- Central tracking of all routes

---

## 📦 API

### `Route.Register(method, path, handler, name, group)`
Registers a route and stores metadata.

```go
core.Route.Register("GET", "/blog", blogHandler, "blog.index", "blog")
```

### `Route.All()`
Returns all registered routes as a slice of `Route` structs.

```go
routes := core.Route.All()
```

### `Route.Get(name)`
Find a route by name.

```go
r, ok := core.Route.Get("blog.index")
fmt.Println(r.Path)
```

### `Route.PathFor(name, ...params)`
Reverse lookup with path params.

```go
url := core.Route.PathFor("blog.view", "slug", "hello-world")
// => "/blog/hello-world"
```

---

## 🧠 How It Works

All routes are stored in an internal registry:

```go
type Route struct {
  Method string
  Path   string
  Name   string
  Group  string
}
```

This registry is populated on startup by calling `core.Route.Register(...)` in each module.

---

## 🛠 Init in `bootstrap.go`

```go
core.InitRouter(echo)
```

This wires the internal Echo instance so `Register()` can attach routes.

---

## 🔥 Use Case Examples

- Automatically generate route docs or admin dashboards
- Enable/disabling routes via config or permissions
- Use `PathFor` in redirects, templates, and link generation

---

## 🛣 Future Enhancements

- Middleware chaining per group
- Versioned route groups (e.g. `/api/v1/...`)
- Aliases or dynamic prefixes

---

[← Back to README](../README.md)

