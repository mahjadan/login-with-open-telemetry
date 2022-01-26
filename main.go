package main

import (
	"github.com/gorilla/mux"
	"github.com/mahjadan/login-with-open-telemetry/cmd/app"
	"github.com/mahjadan/login-with-open-telemetry/cmd/handle"
	"github.com/mahjadan/login-with-open-telemetry/pkg/repository"
	"github.com/mahjadan/login-with-open-telemetry/pkg/service"
	"github.com/mahjadan/login-with-open-telemetry/pkg/token"
	"log"
)

func main() {

	//todo add env variables to be able to change env ( dev, prod)
	repo := repository.NewInMemory()
	tokenMaker := token.NewJWTMaker("my-secret-key")
	srv := service.NewService(repo)
	handler := handle.New(srv, tokenMaker)
	router := mux.NewRouter()

	server := app.New()
	server.SetupRoutes(router, handler)
	log.Println("listening on 8080")
	log.Fatal(server.ListenAndServe())
}
