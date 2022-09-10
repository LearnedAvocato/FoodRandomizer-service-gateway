package main

import (
	"encoding/json"
	"log"
	"net/http"
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
	userData := userData{
		Longitude: ParseFloat(r.URL.Query().Get("longitude")),
		Latitude:  ParseFloat(r.URL.Query().Get("latitude")),
		//UserId:    ParseUint(r.Header.Get("UserId")),
		//HasUserId: len(r.Header.Get("UserId")) > 0,
	}

	payload := jsonRespose{
		Error:    false,
		FoodCard: getRandomFood(userData),
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(jsonPayload)
}
