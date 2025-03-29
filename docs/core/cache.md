# Cache Service (core/cache.go)

The `Cache` service is a lightweight wrapper around [Ristretto](https://github.com/dgraph-io/ristretto), a fast, in-memory key-value cache with support for TTL and cost-based eviction. It also adds support for tag-based cache invalidation.

---

## ✅ Features

- In-memory caching with TTL
- Cost-based eviction policy
- Tag-based cache invalidation
- Thread-safe

---

## 🔧 Initialization

The cache is initialized in `core/bootstrap.go`:

```go
cache := core.NewCache()
```

It is then registered into the `AppServices` struct:

```go
services.Cache.Set("homepage", html, 5*time.Minute)
```

---

## 📦 API

### `Set(key string, value interface{}, ttl time.Duration)`
Store a value in the cache with a TTL.

```go
services.Cache.Set("page:/home", html, 5*time.Minute)
```

### `Get(key string) (interface{}, bool)`
Retrieve a value from the cache.

```go
if html, ok := services.Cache.Get("page:/home"); ok {
    return c.HTML(http.StatusOK, html.(string))
}
```

### `SetTags(key string, tags []string)`
Associate a cache key with one or more tags.

```go
services.Cache.SetTags("page:/home", []string{"tag:homepage"})
```

### `InvalidateTags(tags ...string)`
Delete all keys associated with one or more tags.

```go
services.Cache.InvalidateTags("tag:homepage")
```

---

## 💡 Example: Auto-Caching Rendered Pages

Used with `core.RenderPage()`:

```go
err := core.RenderPage(c, "pages/view", map[string]any{
    "Title": pageTitle,
    "Content": pageContent,
}, true) // Enable caching
```

When caching is enabled, the output is automatically stored and tagged. You can later invalidate related tags when the page is updated:

```go
services.Cache.InvalidateTags("tag:page:" + slug)
```

---

## 🔒 Notes

- Ristretto is **in-memory only** — no persistence between restarts
- TTL granularity is per-entry
- Tag index is managed in memory — efficient, but non-persistent

---

## 🛣 Future Enhancements

- Optional Redis-backed implementation
- Built-in cache stats/inspection route
- Decorator-style helpers (e.g. `WithCache(...)` for handlers)

---

[← Back to README](../README.md)

