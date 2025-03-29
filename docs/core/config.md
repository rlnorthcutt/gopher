# Config (core/config.go)

The `Config` system provides a central place to manage environment variables and application settings, typically loaded at startup via `bootstrap.go`.

It simplifies loading and accessing typed values with sensible defaults.

---

## ✅ Features

- Centralized config values
- Environment variable loading
- Sensible defaults
- Easily extendable for any project

---

## 📦 Config Struct

```go
type Config struct {
    AppName string
    Env     string
    Debug   bool
    Port    string
}
```

You can add more fields depending on your app (e.g. SMTP, API keys, etc).

---

## ⚙️ LoadConfig()

```go
func LoadConfig() *Config {
    return &Config{
        AppName: os.Getenv("APP_NAME"),
        Env:     os.Getenv("APP_ENV"),
        Debug:   os.Getenv("APP_DEBUG") == "true",
        Port:    os.Getenv("PORT"),
    }
}
```

Values are pulled from the environment using `os.Getenv()`.

---

## 🌱 .env File Support

While not required, we recommend using a `.env` file in development with tools like [`godotenv`](https://github.com/joho/godotenv):

```env
APP_NAME=MyCoolApp
APP_ENV=development
APP_DEBUG=true
PORT=8080
```

You can load it at the top of `main.go`:

```go
_ = godotenv.Load()
```

---

## 🛠 Usage

```go
cfg := core.LoadConfig()
fmt.Println(cfg.AppName)
```

---

## 🛣 Future Enhancements

- Support for nested config (database, SMTP, auth, etc)
- JSON/YAML config fallback
- Live reload in development
- CLI overrides (e.g. flags)

---

[← Back to README](../README.md)

