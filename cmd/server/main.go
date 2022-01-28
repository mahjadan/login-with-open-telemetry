package main

import (
	"bytes"
	"context"
	"github.com/gorilla/mux"
	"github.com/mahjadan/login-with-open-telemetry/app/trace"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const appName = "sever"
const environment = "development"

func main() {
	// tracing config
	l := log.New(os.Stdout, "", 0)

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

	log.Println("this is SERVER")
	router := mux.NewRouter()
	router.HandleFunc("/login_test", func(writer http.ResponseWriter, request *http.Request) {
		b := `{ "username": "test1@gmail.com", "password": "123123"}`
		log.Println("going to call ", os.Getenv("LOGIN_URL"))
		resp, err2 := http.DefaultClient.Post(os.Getenv("LOGIN_URL"), "application/json", bytes.NewBuffer([]byte(b)))
		if err2 != nil {
			log.Println(err2)
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err2.Error()))
			return
		}
		if resp.StatusCode == 200 {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(`{"message": "OK"}`))
			return
		} else {
			log.Println("status:", resp.Status)
			writer.WriteHeader(resp.StatusCode)
			all, _ := ioutil.ReadAll(resp.Body)
			writer.Write(all)
		}

	}).Methods(http.MethodGet)

	router.HandleFunc("/register_test", func(writer http.ResponseWriter, request *http.Request) {
		b := `{ "username": "test1@gmail.com", "password": "123123"}`

		log.Println("going to call ", os.Getenv("REGISTER_URL"))
		resp, err2 := http.DefaultClient.Post(os.Getenv("REGISTER_URL"), "application/json", bytes.NewBuffer([]byte(b)))
		if err2 != nil {
			log.Println(err2)
			writer.WriteHeader(http.StatusBadRequest)
			writer.Write([]byte(err2.Error()))
			return
		}
		if resp.StatusCode == 200 {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(`{"message": "OK"}`))
			return
		} else {
			log.Println("status:", resp.Status)
			writer.WriteHeader(resp.StatusCode)
			all, _ := ioutil.ReadAll(resp.Body)
			writer.Write(all)
		}

	}).Methods(http.MethodGet)

	router.Use(otelmux.Middleware("server"))
	port := ":8081"
	log.Println("starting server on ", port)
	server := http.Server{
		Addr:         port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	log.Fatal(server.ListenAndServe())
}
