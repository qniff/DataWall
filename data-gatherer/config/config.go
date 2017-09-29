package config

import (
	"sync"
	"encoding/json"
	"os"
	"fmt"
)

type Configuration struct {
	Url   string
	Token string
}

var conf *Configuration
var once sync.Once

// Return the configuration.
func Get() *Configuration {
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
