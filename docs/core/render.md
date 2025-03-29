# Render Service (core/render.go)

The `Render` system handles HTML output using Go’s `html/template` engine, extended with HTMX support and optional in-memory caching.

---

## ✅ Features

- Go templates (`html/template`)
- HTMX fragment-aware rendering
- Base layout support (`base.html`)
- Helper function map (`safe`, `upper`, etc.)
- Page-level cache with tags + TTL (via `core.Cache`)

---

## 📁 Template Folder Structure

```bash
templates/
├── base.html          # Main layout
├── pages/
│   └── home.html      # Page templates
├── components/        # Reusable blocks
```

---

## 🧱 Core Functions

### `NewTemplateEngine(pattern string)`
Loads and parses templates.

```go
engine, _ := core.NewTemplateEngine("templates/**/*.html")
e.Renderer = engine
```

### `RenderPage(c, templateName, data)`
Renders a full or partial page. Automatically uses `base.html` unless the request is HTMX.

```go
core.RenderPage(c, "pages/home", map[string]any{
    "User": "Alice",
})
```

### `RenderPageNoCache(...)`
Same as `RenderPage`, but skips caching.

### `RenderPageWithCacheFlag(...)`
Internal helper to control cache behavior explicitly.

---

## 📦 HTMX Support

HTMX requests (`HX-Request: true`) return only the template without `base.html` layout.

```go
if isHTMX {
    return c.Render(http.StatusOK, templateName, data)
}
```

---

## 💡 Template Functions

Registered via `DefaultFuncMap()`:

```go
"safe": func(s string) template.HTML {...},
"upper": func(s string) string {...},
```

Add more as needed!

---

## 🔒 Notes

- All templates must be precompiled with `ParseGlob`
- Renders use Echo’s `Renderer` interface
- HTMX detection is automatic

---

## 🛣 Future Enhancements

- Hot reload templates in dev mode
- Custom layout per route
- Component-level cache wrappers

---

[← Back to README](../README.md)

