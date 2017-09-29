package controller

import (
	"DataWall/api/cassandra"
	"encoding/json"
	"fmt"
	"net/http"
)

func getAllLocations(writer http.ResponseWriter, request *http.Request) {
	limit := getLimit(&request.Header)

	data := cassandra.GetData(limit)

	b, err := json.Marshal(data)
	if err != nil {
		fmt.Fprint(writer, DEFAULT_RESPONSE)
	} else {
		fmt.Fprintf(writer, string(b))
	}
}

func root(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.ServeFile(writer, request, "API/api/controller/home.html")
}