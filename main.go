package main

import (
	"flag"
	"log"
	"pipeline/models"
	"pipeline/restaurant"
	"time"
)

var (
	fileName          = "../pipeline/restaurants.csv"
	outputFileFormats = []string{models.JsonData, models.XmlData}
)

func main() {
	start := time.Now()
	log.Println("Reading file: ", fileName)
	sortOption := flag.String("sort", "", "a string")
	ascending := flag.Bool("asc", true, "a string")
	flag.Parse()

	option := restaurants.SortOption{Key: *sortOption, Asc: *ascending}
	restaurantWriter, err := restaurants.NewRestaurantWriter(fileName, outputFileFormats, option)
	if err != nil {
		log.Println("Error in parsing restaurant data", err)
	}
	err = restaurantWriter.WriteRestaurantData()
	if err != nil {
		log.Println("Error in parsing restaurant data", err)
	}

	elapsed := time.Since(start)
	log.Printf("Time taken %s", elapsed)
}
