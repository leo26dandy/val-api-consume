package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// structuring API response here - START
type Response struct {
	Maps []Maps `json:"data"`
}

type Maps struct {
	Name        string   `json:"displayName"`
	Desc        string   `json:"narrativeDescription"`
	MapTactDesc string   `json:"tacticalDescription"`
	MapCoords   string   `json:"coordinates"`
	MapPict     string   `json:"splash"`
	MapPlaces   []Places `json:"callouts"`
}

type Places struct {
	PlaceName   string `json:"regionName"`
	PlaceOnSite string `json:"superRegionName"`
}

// structuring API response here - END

// API Processing here - START
func main() {
	// API to get bundles information
	// response, err := http.Get("https://valorant-api.com/v1/bundles")

	// API to get agents information
	// response, err := http.Get("https://valorant-api.com/v1/agents")

	// API to get maps information
	response, err := http.Get("https://valorant-api.com/v1/maps")

	// check condition
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// API response processing mainly here
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// printing result here
	for _, listMap := range responseObject.Maps {
		fmt.Println("Map Information")
		fmt.Println("\tMap Name:", listMap.Name)
		fmt.Println("\tMap Description:", listMap.Desc)
		fmt.Println("\tMap Tactical Description:", listMap.MapTactDesc)
		fmt.Println("\tMap Coordinates:", listMap.MapCoords)
		fmt.Println("\tMap Picture:", listMap.MapPict)

		fmt.Println("\nMap Callouts")
		for _, listCallouts := range listMap.MapPlaces {
			fmt.Println("\tCallout Name:", listCallouts.PlaceName)
			fmt.Println("\tCallout is On:", listCallouts.PlaceOnSite)
			fmt.Println("\n")
		}

		fmt.Println("---------------------\n")
	}

	// save JSON to local file
	// saveToFile(responseObject)
}

// API Processing here - END

// function to save JSON response into file local
func saveToFile(data Response) {
	// MarshallIndent creates indented JSON representation
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	// Write to a file named "BundleList.json"
	err = os.WriteFile("Maplist.json", jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("List of maps has been saved to MapList.json file")
}
