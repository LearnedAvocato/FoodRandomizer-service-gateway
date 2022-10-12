package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const webPort = "80"

type Config struct {
	conn *grpc.ClientConn
}

func (app *Config) connectToYandexFoodService() {
	conn, err := grpc.Dial("yandex-food-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}
	app.conn = conn
}

func main() {
	app := Config{}

	log.Println("Starting broker service on port", webPort)

	app.connectToYandexFoodService()
	defer app.conn.Close()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
