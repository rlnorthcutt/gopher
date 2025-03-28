package middleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/pocketbase/pocketbase/models"

	"github.com/gopherfw/gopher/core"
)

const ContextUserKey = "currentUser"

// AuthMiddleware verifies a PocketBase token and injects user into the context.
func AuthMiddleware(pb *core.PocketBase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := extractToken(c)
			if token == "" {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "missing auth token",
				})
			}

			user, err := core.VerifyToken(pb, token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{
					"error": "invalid or expired token",
				})
			}

			c.Set(ContextUserKey, user)
			return next(c)
		}
	}
}

// CurrentUser returns the authenticated user from context (or nil).
func CurrentUser(c echo.Context) *models.Record {
	user, ok := c.Get(ContextUserKey).(*models.Record)
	if !ok {
		return nil
	}
	return user
}

// extractToken tries to get token from Bearer header or cookie.
func extractToken(c echo.Context) string {
	// Check Authorization header
	authHeader := c.Request().Header.Get("Authorization")
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}

	// Optional: Check cookie
	if cookie, err := c.Cookie("pb_auth"); err == nil {
		return cookie.Value
	}

	return ""
}
