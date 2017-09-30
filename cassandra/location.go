package cassandra

import "time"

// A wrapper for the database table.
type Device struct {
	X         float32 `json:"x"`
	Y         float32 `json:"y"`
	Z         int     `json:"z"`
	UserType  int     `json:"userType"`
	Hash      string  `json:"hash"`
	CreatedAt time.Time
}
