package config

import (
	"sync"
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
)

// Struct to store all global variables.
type Configuration struct {
	IpAddress string
	Keyspace  string
	ApiPort   int
	Logging   bool
	Token     string
}

var conf *Configuration
var once sync.Once

// Return the configuration.
func Get() *Configuration {
	once.Do(func() {
		absPath, _ := filepath.Abs("../config/config.json")
		file, err := os.Open(absPath)
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&conf)
	})
	return conf
}
