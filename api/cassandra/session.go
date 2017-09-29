package cassandra

import (
	"sync"
	"DataWall/api/config"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
)

var session *gocql.Session
var once sync.Once

func GetSession() *gocql.Session {

	once.Do(func() {
		log.Info("Creating new cassandra instance")
		createSession()
	})
	return session
}

// Create a new session and set the global variable.
func createSession() {
	var err error
	log.Info("Retrieving database configuration")
	var cfg *config.Configuration = config.Get()

	// Connect to the database.
	cluster := gocql.NewCluster(cfg.Ip_addresses)
	cluster.Keyspace = cfg.Keyspace

	// Create the session.
	session, err = cluster.CreateSession()

	// Check for errors.
	if err != nil {
		panic(err)
	}
}
