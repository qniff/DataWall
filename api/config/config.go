package config

import (
	"sync" 								// Config get() func only required to run once
	//"encoding/json" 					// Serializing JSON configuration file to struct
	//log "github.com/sirupsen/logrus" 	// Logging errors
	//"os" 								// Required to open configuration file
)

/**
 * Global configuration struct for app initialization
 */
type Configuration struct {
	Ip_addresses string // IP address of database hosting machine. //TODO List of ip_addresses as cassandra can use clusters.
	Keyspace     string // Name of database
	Port         uint16 // Port is always positive and below 65535
}

var conf *Configuration // Predeclared global configuration struct, accessible by Get() func returning conf pointer.

/** Config.Get
 * Opens json config file. Decodes content to json into var conf of Configuration struct.
 * @return Configuration pointer
 */
func Get() *Configuration {
	sync.Once.Do(func() {
		/*var configPath string = "config/config.json"
		logrus.Info("Retrieving and setting configuration from file: ")
		//TODO Dynamic configuration from json config file.
		//BUG Below code works except for IP field, crashes as is nil. Change ip_addresses field in configuration struct back to Array.
		file, fileErr := os.Open(configPath)
		defer file.Close()

		if fileErr != nil {
			log.WithFields(log.Fields{
				"error": fileErr.Error(),
			}).Fatal("Config file error")
		}

		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&conf)*/

		conf = &Configuration{
			Ip_addresses: "192.168.31.131",
			Keyspace:     "data",
			Port:         8081,
		}
	})
	return conf
}
