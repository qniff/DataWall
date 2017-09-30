package controller

import (
	"DataWall/cassandra"
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllLocations(writer http.ResponseWriter, request *http.Request) {
	limit := getLimit(&request.Header)

	data := cassandra.GetDevices(limit)

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprint(writer, defaultResponse)
	} else {
		fmt.Fprintf(writer, string(b))
	}
}
