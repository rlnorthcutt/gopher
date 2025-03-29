# Job Scheduler (core/jobs.go)

The `JobScheduler` provides a simple but powerful API for running scheduled background tasks using the [gocron](https://github.com/go-co-op/gocron) library.

Jobs can be registered by modules and will automatically run on a defined interval or cron expression.

---

## ✅ Features

- Register recurring jobs with cron expressions
- Automatically injects `CoreServices` into jobs
- Built-in logging (start, finish, panic recovery)
- Run jobs on a schedule or on-demand by name

---

## 🧱 Usage Example

```go
func ClearCache(s *core.CoreServices) {
    s.Cache.InvalidateTags("tag:homepage")
    s.Logger.Info().Msg("Homepage cache cleared")
}

func (m *BlogModule) RegisterJobs(j *core.JobScheduler) {
    j.Register("clear_homepage_cache", "@every 5m", ClearCache)
}
```

---

## 📦 API

### `Register(name, cronExpr, handler)`
Registers a job using cron syntax.

```go
j.Register("cleanup_tokens", "@every 30m", Cleanup)
```

### `Start()`
Begins running scheduled jobs asynchronously.

```go
services.Jobs.Start()
```

### `RunNow(name)`
Executes a named job manually.

```go
services.Jobs.RunNow("clear_homepage_cache")
```

---

## 💡 Logging & Recovery

Each job automatically logs:
- Start and completion
- Duration
- Panic recovery (with error logging)

You do not need to manually log these unless you want more context.

---

## 🔐 Access to Core Services

Each job handler receives a pointer to the app’s `CoreServices`:

```go
func Cleanup(s *core.CoreServices) {
    s.Logger.Info().Msg("Cleaning up...")
}
```

---

## 🛣 Future Enhancements

- CLI: list, run, enable/disable jobs
- DB-persisted jobs (editable from admin)
- Retry/failure tracking and reporting
- Timezone-aware scheduling

---

[← Back to README](../README.md)

