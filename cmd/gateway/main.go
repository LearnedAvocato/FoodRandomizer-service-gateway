package main

import (
	"log"
	"net/http"

	"gateway/internal/app/gateway"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.Println("Starting gateway service on port 80")

	conn, err := grpc.Dial("yandex-food-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}
	defer conn.Close()

	gateway := gateway.NewGateway(conn)

	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST"},
	}))
	mux.Use(middleware.Heartbeat("/ping"))
	mux.Get("/getRandomFood", gateway.GetRandomFood)
	srv := &http.Server{
		Addr:    "0.0.0.0:80",
		Handler: mux,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
