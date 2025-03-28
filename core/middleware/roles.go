package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RequireAuth ensures a user is present in the context.
func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := CurrentUser(c)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"error": "authentication required",
			})
		}
		return next(c)
	}
}

// RequireRole ensures the current user has a specific role.
func RequireRole(role string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := CurrentUser(c)
			if user == nil || user.GetString("role") != role {
				return c.JSON(http.StatusForbidden, echo.Map{
					"error": "forbidden: role required",
				})
			}
			return next(c)
		}
	}
}

// RequireRoles allows access if the user has ANY of the listed roles.
func RequireRoles(allowed []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := CurrentUser(c)
			if user == nil {
				return c.JSON(http.StatusUnauthorized, echo.Map{"error": "unauthenticated"})
			}

			userRole := user.GetString("role")
			for _, role := range allowed {
				if role == userRole {
					return next(c)
				}
			}
			return c.JSON(http.StatusForbidden, echo.Map{"error": "forbidden"})
		}
	}
}

// RequireCondition lets you provide a custom access check.
func RequireCondition(check func(*echo.Context) bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !check(&c) {
				return c.JSON(http.StatusForbidden, echo.Map{
					"error": "access denied",
				})
			}
			return next(c)
		}
	}
}
