package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sherinur/soar-system/backend/auth_service/config"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	otelProm "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.uber.org/zap"
)

const telemetryIPAddress = "0.0.0.0:%d"

type Telemetry struct {
	TracerProvider *sdktrace.TracerProvider
	MetricProvider *metric.MeterProvider

	cfg *config.Telemetry
	log *zap.Logger
}

func InitTelemetry(ctx context.Context, cfg config.Telemetry, log *zap.Logger) (*Telemetry, error) {
	// open telemetry
	traceExp, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint("localhost:4318"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	// tracer provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	// prometheus
	metricExp, err := otelProm.New()
	if err != nil {
		return nil, err
	}

	// metric provider
	mp := metric.NewMeterProvider(
		metric.WithReader(metricExp),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
		)),
	)

	// metric exporter endpoint
	go serveMetrics(cfg.ExporterPromPort, log)

	otel.SetTracerProvider(tp)
	otel.SetMeterProvider(mp)

	return &Telemetry{
		TracerProvider: tp,
		MetricProvider: mp,
		log:            log,
	}, nil
}

func serveMetrics(port int, log *zap.Logger) {
	mux := http.NewServeMux()

	addr := fmt.Sprintf(telemetryIPAddress, port)

	mux.Handle("/metrics", promhttp.Handler())

	log.Info("Starting serving metrics on http", zap.String("addr", addr+"/metrics"))
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Error("Metrics server error", zap.Error(err))
	}
}
