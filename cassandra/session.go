package cassandra

import (
	"sync"

	"DataWall/config"

	log "github.com/sirupsen/logrus" // Logging library
	"github.com/gocql/gocql"         // Cassandra database driver
)

var _session *gocql.Session
var once sync.Once


/** createSession
 & Serves to only call createSession func only once otherwise referred to already set global _session variable.
 */
func session() *gocql.Session {
	// Block executes only once
	once.Do(func() {
		log.Info("Creating new cassandra instance")
		createSession()
	})
	return _session
}

/** createSession
 & Instantiate new cassandra database session. Sets predeclared global _session variable.
 */
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
