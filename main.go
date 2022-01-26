package main

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mahjadan/login-with-open-telemetry/cmd/app"
	"github.com/mahjadan/login-with-open-telemetry/cmd/handle"
	"github.com/mahjadan/login-with-open-telemetry/cmd/trace"
	"github.com/mahjadan/login-with-open-telemetry/pkg/repository"
	"github.com/mahjadan/login-with-open-telemetry/pkg/service"
	"github.com/mahjadan/login-with-open-telemetry/pkg/token"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"log"
	"os"
)

/*
docker run -d --name jaeger \
  //-e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.30
*/
const appName = "login-service"
const environment = "development"

func main() {

	// tracing config
	l := log.New(os.Stdout, "", 0)

	//exp, err := trace.NewStdoutExporter(os.Stdout) // this is for stdout exporter
	//url := "http://localhost:14268/api/traces" // this is for http endpoint exporter.
	//exp, err := trace.NewJaegerExporterHttp(url)
	exp, err := trace.NewJaegerExporterGrpc()
	if err != nil {
		l.Fatal(err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(trace.NewResource(appName, environment)),
	)
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			l.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	//todo add env variables to be able to change env ( dev, prod)
	repo := repository.NewInMemory()
	tokenMaker := token.NewJWTMaker("my-secret-key")
	srv := service.NewService(repo, appName)
	handler := handle.New(srv, tokenMaker)
	router := mux.NewRouter()

	router.Use(otelmux.Middleware(appName))

	server := app.New("8080")
	server.SetupRoutes(router, handler)
	log.Println("listening on 8080")
	log.Fatal(server.ListenAndServe())
}
