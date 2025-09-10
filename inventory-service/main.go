package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Order Service is good!")
	})

	log.Println("Starting Order Service on port 8082...")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatal(err)
	}
}