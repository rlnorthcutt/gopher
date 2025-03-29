# Permissions (core/permissions.go)

The `PermissionsManager` provides a programmatic way to manage PocketBase collection access rules (list, view, create, update, delete).

It also supports loading rules from external YAML files, making it easy to manage collection security declaratively.

---

## ✅ Features

- Define collection permissions in Go
- Sync rules with PocketBase on startup or migration
- YAML support for config-based rule sets
- Avoids manual editing in PB Admin UI

---

## 🧱 RuleSet Struct

```go
type RuleSet struct {
    List   *string
    View   *string
    Create *string
    Update *string
    Delete *string
}
```

- `nil`: skip rule (keep existing)
- `""`: clear rule (public access)
- `"..."`: apply rule string

---

## 🧩 Set Rules in Code

```go
services.Permissions.SetRules("posts", core.RuleSet{
    List:   ptr("@request.auth.id != ''"),
    View:   ptr("@request.auth.id != ''"),
    Update: ptr("@collection.posts.author = @request.auth.id"),
})
```

Use `ptr()` helper to simplify string pointers:

```go
func ptr(s string) *string {
    return &s
}
```

---

## 📄 Load from YAML

Create a `permissions.yaml` file in a module:

```yaml
collection: posts
rules:
  list: "@request.auth.id != ''"
  view: "@request.auth.id != ''"
  update: "@collection.posts.author = @request.auth.id"
```

Then load and apply:

```go
err := services.Permissions.LoadFromYAML("modules/blog/permissions.yaml")
```

---

## 🔒 Best Practices

- Use `SetRules()` in `Migrate()` for each module
- Use YAML if you want editable configs or external review
- Keep rules scoped and auditable — never hardcode logic in views

---

## 🛣 Future Enhancements

- CLI: list, show, set, reset rules
- Describe rules in YAML for UI editors
- UI for syncing and previewing permissions

---

[← Back to README](../README.md)