# Entity System (core/entity.go)

The Entity system provides a generic way to define Go structs as domain models while leveraging PocketBase for storage.

It supports saving, updating, deleting, and retrieving data through the `core.DataService`, and enables you to hook into lifecycle events.

---

## ✅ Features

- Define reusable models with Go structs
- Automatic Save / Delete / Update routing
- Hook support (before/after save/delete)
- Works with PocketBase collections

---

## 📦 Interface

### `Entity`

```go
type Entity interface {
    Collection() string
    ID() string
    SetID(string)
    Load(record *models.Record)
    ToRecord() *models.Record
}
```

Implement this on any struct to make it a valid entity.

---

## 🧩 Example Entity

```go
type Page struct {
    ID      string
    Title   string
    Slug    string
    Content string
}

func (p *Page) Collection() string   { return "pages" }
func (p *Page) ID() string           { return p.ID }
func (p *Page) SetID(id string)      { p.ID = id }
func (p *Page) Load(r *models.Record) {
    p.ID = r.Id
    p.Title = r.GetString("title")
    p.Slug = r.GetString("slug")
    p.Content = r.GetString("content")
}
func (p *Page) ToRecord() *models.Record {
    r := models.NewRecord("pages")
    r.Set("title", p.Title)
    r.Set("slug", p.Slug)
    r.Set("content", p.Content)
    return r
}
```

---

## 💾 Save or Delete

With an active `CoreServices` instance:

```go
page := &Page{Title: "About", Slug: "about", Content: "Hello!"}
err := services.Data.Save(page)
```

Delete:

```go
services.Data.Delete(page)
```

---

## 🔄 Hooks (WIP)

Future versions will support:
- `BeforeSave()` / `AfterSave()`
- `BeforeDelete()` / `AfterDelete()`

And allow for both interface-based and registry-based hooks.

---

## 🛣 Future Enhancements

- Auto-generate PocketBase schema from entity structs
- Validation helpers (e.g. required fields, regex)
- Type conversion helpers for records
- Integration with modules

---

[← Back to README](../README.md)

