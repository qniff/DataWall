package api

import (
	"fmt"
	"log"
	"net/http"
	"DataWall/API/config"
	"DataWall/API/api/controller"
)

// Start serving the api.
func Start() {
	controller.RegisterEndPoints()

	log.Printf("Running at: %s:%d\n\n", config.Get().Ip_addresses, config.Get().Port)

	http.ListenAndServe(fmt.Sprint(":", config.Get().Port), nil)
}