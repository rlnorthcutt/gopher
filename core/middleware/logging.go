package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
)

func LoggingMiddleware(logger zerolog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			req := c.Request()
			res := c.Response()

			event := logger.Info().
				Str("method", req.Method).
				Str("path", req.URL.Path).
				Int("status", res.Status).
				Dur("latency", stop.Sub(start)).
				Str("userAgent", req.UserAgent()).
				Str("remoteIP", c.RealIP()).
				Str("request_id", GetRequestID(c)) // ✅ Add request_id

			// Optional: attach trace ID if available
			if spanCtx := trace.SpanContextFromContext(req.Context()); spanCtx.IsValid() {
				event = event.Str("trace_id", spanCtx.TraceID().String())
			}
			event.Msg("request")

			return err
		}
	}
}
