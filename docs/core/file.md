# File Service (core/file.go)

The File service provides a simple abstraction over PocketBase's file storage system. It enables uploading, retrieving, and deleting files, and provides helpers to generate public URLs.

---

## ✅ Features

- Upload files to any PB collection/record
- Get secure or public access URLs
- Delete files from PB
- Future expansion for file metadata + CDN

---

## 📦 API

### `Upload(collection, recordID, field, file)`
Uploads a file to the specified field.

```go
url, err := services.File.Upload("posts", record.Id, "image", file)
```

### `GetURL(collection, recordID, fileName)`
Returns the public URL for a file.

```go
url := services.File.GetURL("posts", record.Id, "header.png")
```

### `Delete(collection, recordID, fileName)`
Deletes a file from storage.

```go
err := services.File.Delete("posts", record.Id, "old.png")
```

---

## 🔒 Access + Permissions

- File visibility and access are based on the PocketBase collection’s rules.
- Private file access may require signed URLs (future feature).

---

## 📁 File Organization

Files are stored in PocketBase’s internal file storage by collection/record/field. This is automatic.

---

## 🛠 Use in Handlers

```go
file, _ := c.FormFile("upload")
url, err := services.File.Upload("pages", record.Id, "attachment", file)
```

---

## 🛣 Future Enhancements

- Support for multiple file fields (uploads[])
- File metadata access (size, mime, etc.)
- Signed URLs or token-based access
- CDN/cache-layer integration

---

[← Back to README](../README.md)

