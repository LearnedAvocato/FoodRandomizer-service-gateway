package main

import (
	"context"
	proto "gateway/proto/generated"
	"log"
	"net/http"
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

func getRandomFood(userData userData) foodCard {
	// TODO: check user id
	log.Print("User coordinates: ", userData.Latitude, userData.Longitude)

	return foodCard{
		ImageUrl:    "url",
		Description: "tasty",
		Title:       "food",
	}
}

func (app *Config) GetRandomFood(w http.ResponseWriter, r *http.Request) {
	log.Println("get food request")
	log.Println("1")
	conn, err := grpc.Dial("delivery-club-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	log.Println("2t")
	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}
	defer conn.Close()

	c := proto.NewDeliveryClubClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("do food request")
	foodResponse, err := c.GetRandomFood(ctx, &proto.FoodRequest{
		CardsNum:   1,
		FoodFilter: &proto.FoodFilter{},
		Longitude:  37.762069,
		Latitude:   55.669746,
	})

	log.Println("foodResponse received")

	if err != nil {
		log.Fatalf("Failed to get random food by gRPC: %v", err)
	}

	jsonPayload := protojson.Format(foodResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(jsonPayload))
}
