package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		data := []byte(`{"pong"}`)
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)

}
