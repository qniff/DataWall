package main

import (
	"encoding/json"
	"fmt"
	"time"
	"io/ioutil"

	"golang.org/x/oauth2"
	"log"
	"DataWall/cassandra"
	"DataWall/config"
)

const devicesEndpointUrl string = "https://api.fhict.nl/location/devices"
const interval time.Duration = 20000 * time.Millisecond

var logging bool
var token string

func main() {
	fmt.Println("Application started...")

	logging = config.Get().Logging
	token = config.Get().Token

	// Start the function call loop
	doEvery(interval, getDataFromApi)
}

func doEvery(interval time.Duration, repeatFunction func(time.Time)) {
	for currentTime := range time.Tick(interval) {
		repeatFunction(currentTime)
	}
}

func getDataFromApi(currentTime time.Time) {
	if logging {
		fmt.Println("Start:", currentTime)
	}

	// Get a response from the Fontys API.
	tokenSource := &TokenSource{
		AccessToken: token,
	}
	resp, _ := oauth2.NewClient(oauth2.NoContext, tokenSource).Get(devicesEndpointUrl)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	jsonData := string(body)

	// Convert the JSON to a struct array.
	var devices []cassandra.Device
	err := json.Unmarshal([]byte(jsonData), &devices)
	if err != nil {
		log.Panic("Could not convert json data to struct. Data from Fontys API:", jsonData)
	}

	// Store the data in the database.
	cassandra.InsertDevices(devices)

	if logging {
		fmt.Println("End:", time.Now())
	}
}
