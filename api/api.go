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

	cfg := *config.Get()
	log.Printf("API running at: %s:%d\n\n", cfg.IpAddress, cfg.ApiPort)

	http.ListenAndServe(fmt.Sprint(":", cfg.ApiPort), nil)
}
