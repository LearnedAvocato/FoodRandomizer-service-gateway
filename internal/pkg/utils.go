package main

import (
	"log"
	"strconv"
)

func ParseFloat(str string) float64 {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Println(err)
	}
	return num
}

func ParseUint(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return num
}
