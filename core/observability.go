package core

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

// SetupLogger configures zerolog for structured logging.
func SetupLogger(serviceName string) zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", serviceName).
		Logger()

	log.Logger = logger // global override for `log.Print` etc.
	return logger
}

// SetupTracer initializes OpenTelemetry tracer with basic stdout exporter.
func SetupTracer(serviceName string) (otel.Tracer, error) {
	exporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	otel.SetTracerProvider(tp)

	return otel.Tracer(serviceName), nil
}

// AttachMetricsRoute adds a basic Prometheus metrics endpoint.
func AttachMetricsRoute(e *echo.Echo) {
	e.GET("/metrics", func(c echo.Context) error {
		return c.String(200, "# TODO: integrate Prometheus exporter\n")
	})
}
