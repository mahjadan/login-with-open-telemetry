module github.com/mahjadan/login-with-open-telemetry

go 1.17

require (
	github.com/golang-jwt/jwt/v4 v4.2.0
	github.com/gorilla/mux v1.8.0
	github.com/pkg/errors v0.9.1
	github.com/satori/go.uuid v1.2.0
	go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux v0.28.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.28.0
	go.opentelemetry.io/otel v1.3.0
	go.opentelemetry.io/otel/exporters/jaeger v1.3.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.3.0
	go.opentelemetry.io/otel/sdk v1.3.0
	golang.org/x/crypto v0.0.0-20220126173729-e04a8579fee6
)

require (
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logr/logr v1.2.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/otel/internal/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/metric v0.26.0 // indirect
	go.opentelemetry.io/otel/trace v1.3.0 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
