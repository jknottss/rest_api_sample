package main

import (
	rest_api "REST_API"
	"REST_API/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(rest_api.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running http server: %s\n", err.Error())
	}
}
