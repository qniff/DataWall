package controller

import (
	"net/http"
	"strconv"
)

const DEFAULT_LIMIT = 50
const HEADER_LIMIT = "Limit"
const DEFAULT_RESPONSE = "{}"

func RegisterEndPoints() {
	// GET: /
	http.HandleFunc("/", root)
	// GET: /data
	http.HandleFunc("/data", getAllLocations)
}

// Return the value supplied by the "Limit" header.
// If the value is not an integer return the default value.
func getLimit(h *http.Header) int {
	tempLimit := h.Get(HEADER_LIMIT)
	limit, err := strconv.Atoi(tempLimit)
	if err != nil {
		return DEFAULT_LIMIT
	}
	return limit
}

