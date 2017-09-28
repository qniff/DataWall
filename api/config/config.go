package config

import (
	"sync"
	//"encoding/json"
	//"os"
	//"fmt"
	//log "github.com/sirupsen/logrus"
)

// Struct to store all global variables.
type Configuration struct {
	Ip_addresses string
	Keyspace string
	Port int
}

var conf *Configuration
var once sync.Once

// TODO: Move settings to a json/yaml/xml/gcfg file.
// Return the configuration.
func Get() *Configuration{
	once.Do(func() {
		/*file, err := os.Open("config/config.json")
		defer file.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&conf)*/

		conf = &Configuration{
			Ip_addresses:"192.168.31.131",
			Keyspace: "data",
			Port: 8081,
		}
	})
	return conf
}
