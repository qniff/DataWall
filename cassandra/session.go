package cassandra

import (
	"sync"
	"github.com/gocql/gocql"
	log "github.com/sirupsen/logrus"
	"DataWall/config"
)

var _session *gocql.Session
var once sync.Once

func session() *gocql.Session {

	once.Do(func() {
		log.Info("Creating new cassandra instance")
		createSession()
	})
	return _session
}

// Create a new _session and set the global variable.
func createSession() {
	var err error

	cfg := *config.Get()
	// Connect to the database.
	cluster := gocql.NewCluster(cfg.IpAddress)
	cluster.Keyspace = cfg.Keyspace

	// Create the _session.
	_session, err = cluster.CreateSession()

	// Check for errors.
	if err != nil {
		panic(err)
	}
}