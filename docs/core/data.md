# Data Service (core/data.go)

The `DataService` provides a convenience layer over PocketBase's data access system. It enables you to fetch, filter, and work with records easily, using a consistent Go-friendly API.

---

## ✅ Features

- Find records with filters or IDs
- Convert to/from Go structs (`Entity`)
- First/All lookup helpers
- Fully typed with support for extension

---

## 📦 API Overview

### `FindFirst(collection, filter) → *Record`
Finds the first record that matches the given filter map.

```go
record, err := services.Data.FindFirst("pages", map[string]any{"slug": "about"})
```

### `FindAll(collection, filter) → []*Record`
Finds all matching records.

```go
posts, _ := services.Data.FindAll("posts", map[string]any{"author": userId})
```

### `FindByID(collection, id)`
Lookup by ID.

```go
record, err := services.Data.FindByID("users", id)
```

---

## 🔄 Save / Delete via `Entity`

You can pass any object that implements the `Entity` interface:

```go
err := services.Data.Save(myPost)
err := services.Data.Delete(myPost)
```

This automatically converts to/from `*models.Record` and performs persistence using the PocketBase DAO.

---

## ⚙️ Internal Helpers

- `toFilterExpr(map[string]any)`
  Converts a simple Go map to PocketBase filter expressions.
- `recordToMap(*models.Record)`
  Extracts field values from a record.

---

## 🛣 Future Enhancements

- Soft delete support
- Pagination + sorting helpers
- Query builder DSL
- Auto-mapping to entity structs via reflection or struct tags

---

[← Back to README](../README.md)