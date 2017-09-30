package controller

import (
	"net/http"
	"strconv"
	"path/filepath"
)

const defaultLimit = 50
const headerLimit = "Limit"
const defaultResponse = "{}"

func RegisterEndPoints() {
	// GET: /
	http.HandleFunc("/", root)
	// GET: /data
	http.HandleFunc("/data", getAllLocations)
}

// Return the value supplied by the "Limit" header.
// If the value is not an integer return the default value.
func getLimit(h *http.Header) int {
	tempLimit := h.Get(headerLimit)
	limit, err := strconv.Atoi(tempLimit)
	if err != nil {
		return defaultLimit
	}
	return limit
}

// The default root which serves the static file with the endpoints.
func root(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	absPath, _ := filepath.Abs("../api/static/home.html")
	http.ServeFile(writer, request, absPath)
}
