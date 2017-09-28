package cassandra

import "time"

// A wrapper for the database table.
type Location struct {
	LocX      float32
	LocY      float32
	LocZ      int
	UserHash  string
	CreatedAt time.Time
}
