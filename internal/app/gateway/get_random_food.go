package gateway

import (
	"context"
	yandex_food "gateway/pkg/api/yandex-food"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/protobuf/jsonpb"
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
}

func (i *Implementation) GetRandomFood(w http.ResponseWriter, r *http.Request) {
	yandexFoodClient := yandex_food.NewYandexFoodClient(i.yandexFoodConn)

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
	cardsNum, err := strconv.ParseInt(r.URL.Query().Get("cardsNum"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to get latitude from request: %v", err)
	}
	getTags := false
	getTagsStr := r.URL.Query().Get("getTags")
	if getTagsStr == "true" {
		getTags = true
	}

	selectedTags := r.URL.Query()["tags"]

	foodResponse, err := yandexFoodClient.GetRandomFood(ctx, &yandex_food.FoodRequest{
		CardsNum:     cardsNum,
		GetTags:      getTags,
		Longitude:    float32(longitude),
		Latitude:     float32(latitude),
		SelectedTags: selectedTags,
	})

	if err != nil {
		log.Printf("Failed to get random food by gRPC: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	marshaller := jsonpb.Marshaler{}
	jsonPayload, err := marshaller.MarshalToString(foodResponse)
	if err != nil {
		log.Printf("Error while converting json to string: %v", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(jsonPayload))
}
