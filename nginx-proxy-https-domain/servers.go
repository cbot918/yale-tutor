package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	port1 = ":3000"
	port2 = ":3001"
)

func main() {

	go server(port1)
	go server(port2)
	fmt.Println("starting server cluster")

	select {}
}

func server(port string) {
	mux := http.NewServeMux()

	// Register a handler function to the new mux
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server: %s\n", port)
	})

	fmt.Printf("Listening on %s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Failed to start server on %s: %v", port, err)
	}
}
