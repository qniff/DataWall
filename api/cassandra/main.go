package cassandra

import (
	"github.com/gocql/gocql"
	"time"
	"fmt"
)

// Retrieve data from the database.
// limit: The amount of rows you want to limit your request to.
func GetData(limit int) []Location {
	var s *gocql.Session = GetSession()
	var locList []Location
	m := map[string]interface{}{}

	// Create the query.
	query := fmt.Sprintf(
		"%s %d",
		"SELECT user_hash, createdat, loc_x, loc_y, loc_z FROM locations LIMIT ",
		limit)
	iterable := s.Query(query).Iter()

	// Iterate over results.
	for iterable.MapScan(m) {
		locList = append(locList, Location{
			UserHash:  m["user_hash"].(string),
			CreatedAt: m["createdat"].(time.Time),
			LocX:      m["loc_x"].(float32),
			LocY:      m["loc_y"].(float32),
			LocZ:      m["loc_z"].(int),
		})
		m = map[string]interface{}{}
	}

	return locList
}

