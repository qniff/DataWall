package cassandra

import (
	"time"
	"fmt"
	"log"
)

// Retrieve data from the database.
// limit: The amount of rows you want to limit your request to.
func GetDevices(limit int) []Device {
	var locList []Device
	m := map[string]interface{}{}

	// Create the query.
	query := fmt.Sprintf(
		"%s %d",
		"SELECT user_hash, createdat, loc_x, loc_y, loc_z FROM locations LIMIT ",
		limit)
	iterable := session().Query(query).Iter()

	// Iterate over results.
	for iterable.MapScan(m) {
		locList = append(locList, Device{
			Hash:      m["user_hash"].(string),
			CreatedAt: m["createdat"].(time.Time),
			X:         m["loc_x"].(float32),
			Y:         m["loc_y"].(float32),
			Z:         m["loc_z"].(int),
		})
		m = map[string]interface{}{}
	}

	return locList
}

func InsertDevices(devices []Device) {
	for _, device := range devices  {
		go insert(device)
	}
}

func insert(device Device) {
	if err := session().Query(`INSERT INTO locations (loc_x, loc_y, loc_z, user_hash, createdat) VALUES (?, ?, ?, ?, ?)`,
		device.X, device.Y, device.Z, device.Hash, time.Now()).Exec(); err != nil {
		log.Fatal(err)
	}
}
