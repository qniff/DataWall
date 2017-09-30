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
	var sessionErr error

	// Retrieve configuration for DB IP and Keyspace.
	cfg := *config.Get()
	// Connect to cassandra database cluster and databas keyspace (DB name).
	cluster := gocql.NewCluster(cfg.IpAddress)
	cluster.Keyspace = cfg.Keyspace

	// Set global session variable
	_session, sessionErr = cluster.CreateSession()
	// Handling possible create session error
	if sessionErr != nil {
		// Logging & entering Panic state.
		log.WithFields(log.Fields{
			"Error": sessionErr.Error(),
		}).Panic("Failed to create Cassandra DB session!")
		// TODO Recover from panic?
	}
}
