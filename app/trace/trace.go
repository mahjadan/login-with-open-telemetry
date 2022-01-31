package trace

import (
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.7.0"
	"io"
	"net/http"
)

func NewJaegerExporterGrpc() (trace.SpanExporter, error) {
	// this required OTEL_EXPORTER_JAEGER_AGENT_HOST and
	// OTEL_EXPORTER_JAEGER_AGENT_PORT
	//or you can pass the values using jaeger.WithAgentHost() , jaeger.WithAgentPort()
	exp, err := jaeger.New(jaeger.WithAgentEndpoint())
	if err != nil {
		return nil, err
	}
	return exp, nil
}

func HTTPClientWithTrace() *http.Client {
	client := http.DefaultClient
	transport := otelhttp.NewTransport(client.Transport)
	client.Transport = transport
	return client
}
func NewJaegerExporterHttp(url string) (trace.SpanExporter, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return exp, nil
}
func NewStdoutExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

func NewResource(appName, environment string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(appName),
			semconv.ServiceVersionKey.String("v0.1.0"),
			attribute.String("environment", environment),
		),
	)
	return r
}
