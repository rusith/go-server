package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHello)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := strings.TrimPrefix(r.URL.Path, "/")
	message = "Hello " + message + "!"
	w.Header().Set("Content-Type", "application/json")
	jsonD, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		w.Write([]byte("\"An error occurred while trying to process the request\""))
		w.WriteHeader(500)
	}
	w.Write(jsonD)
}