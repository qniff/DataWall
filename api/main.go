package main

import (
	"DataWall/api/cassandra"
	log "github.com/sirupsen/logrus"
	"DataWall/api/api"
)

// Questionable file, future tbd.
func main() {
	testing := false
	log.SetLevel(log.InfoLevel)

	if testing {
		log.SetLevel(log.InfoLevel)
		log.WithFields(log.Fields{
			"testing": testing,
		}).Info("Testing environmental variable is set!")
		cassandra.RunTests()
	} else {
		api.Start()
	}
}
