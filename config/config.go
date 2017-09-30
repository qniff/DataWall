package config

import (
	"sync"
	"fmt"
	"encoding/json"
	"os"
)

// Struct to store all global variables.
type configuration struct {
	IpAddress string
	Keyspace  string
	ApiPort   int
	Logging   bool
	Token     string
}

var conf *configuration
var once sync.Once

// Return the configuration.
func Get() *configuration {
	once.Do(func() {
		file, err := os.Open("config/config.json")
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&conf)
	})
	return conf
}
