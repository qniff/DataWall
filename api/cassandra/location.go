package cassandra

import "time"

/**
 * Struct for Fontys API return result
 */
type Location struct {
	LocX      float32	// X axis of device location. Generated from single geolocation.
	LocY      float32	// Y axis of device location. Generated from single geolocation.
	LocZ      int8		// Z-index determines floor with one digit. Not making unsigned as future might make basement negative?
	UserHash  string	// UserHash from fontys is a combination of Ip address and datetime (presumed of IP handout).
	CreatedAt time.Time // CreatedAt field is added and assigned instantly at return
}