package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const ContextRequestID = "request_id"

// RequestIDMiddleware generates a request ID and injects it into context and headers.
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			id := uuid.NewString()

			c.Set(ContextRequestID, id)
			c.Response().Header().Set("X-Request-ID", id)

			return next(c)
		}
	}
}

// GetRequestID returns the request ID from context.
func GetRequestID(c echo.Context) string {
	if val, ok := c.Get(ContextRequestID).(string); ok {
		return val
	}
	return ""
}
