package config

import (
	"sync"          // Config get() func only required to run once
	"encoding/json" // Serializing JSON configuration file to struct
	"os"            // Open configuration file
	"path/filepath" // Define config file path

	log "github.com/sirupsen/logrus" // Logging errors
)

// Struct to store all global variables.
type Configuration struct {
	IpAddress string // IP addresses of cassandra database
	Keyspace  string // Keyspace (DB name) of cassandra database
	ApiPort   uint8  // The port on which the API will run. Always positive and below 65535
	Logging   bool   // Whether to display logs or not
	Token     string // Auth token from Fontys API
}

var conf *Configuration // Predeclared global configuration struct, accessible by Get() func returning conf pointer.

/** Config.Get
 * Opens json config file. Decodes content to json into var conf of Configuration struct.
 * @return Configuration pointer
 */
func Get() *Configuration {
	sync.Once.Do(func() {
		// Set file path to be serialized
		absPath, _ := filepath.Abs("../config/config.json")

		// All logs by contextLogger now include file path
		contextLogger := log.WithFields(log.Fields{
			"path": absPath,
		})
		contextLogger.Info("Retrieving and setting configuration from file...")
		file, fileErr := os.Open(absPath)
		defer file.Close()
		if fileErr != nil {
			log.WithFields(log.Fields{
				"Error": fileErr.Error(),
			}).Fatal("Failed opening file! ")
		}

		// Decoding (serializing) JSON content of file to Configuration struct
		jsonParser := json.NewDecoder(file)
		jsonParser.Decode(&conf)
		contextLogger.Info("Configuration initialized!")
	})
	return conf
}
