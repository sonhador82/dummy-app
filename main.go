package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8000

func main() {
	log.Println("Start listening:", port)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"status\": \"ok\"}"))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
