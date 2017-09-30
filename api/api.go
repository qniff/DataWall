package main

import (
	"fmt"
	"log"
	"net/http"
	"DataWall/config"
	"DataWall/api/controller"
)

// Start serving the api.
func main() {
	controller.RegisterEndPoints()

	log.Printf("API running at: %s:%d\n\n", config.Get().IpAddress, config.Get().Port)

	http.ListenAndServe(fmt.Sprint(":", config.Get().Port), nil)
}
