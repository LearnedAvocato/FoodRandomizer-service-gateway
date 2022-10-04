package main

import (
	"context"
	proto "gateway/proto/generated"
	"log"
	"net/http"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/encoding/protojson"
)

type foodCard struct {
	ImageUrl    string `json:"imageUrl"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

type jsonRespose struct {
	Error    bool     `json:"error"`
	FoodCard foodCard `json:"food"`
}

type userData struct {
	Longitude float64
	Latitude  float64
	//UserId    uint64
	//HasUserId bool
}

// CONNECT ONLY 1 TIME
func (app *Config) GetRandomFood(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial("yandex-food-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}
	defer conn.Close()

	c := proto.NewYandexFoodClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	longitude, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 32)
	if err != nil {
		log.Fatalf("Failed to get longitude from request: %v", err)
	}
	latitude, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 32)
	if err != nil {
		log.Fatalf("Failed to get latitude from request: %v", err)
	}

	foodResponse, err := c.GetRandomFood(ctx, &proto.FoodRequest{
		CardsNum:   8,
		FoodFilter: &proto.FoodFilter{},
		Longitude:  float32(longitude),
		Latitude:   float32(latitude),
	})

	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}

	jsonPayload := protojson.Format(foodResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(jsonPayload))
}
