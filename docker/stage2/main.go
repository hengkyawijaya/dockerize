package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	ServiceName = os.Getenv("SERVICE_NAME")
)

func main() {
	log.Printf("server %s started\n", ServiceName)
	http.HandleFunc("/stage2", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "stage2")
	})

	http.ListenAndServe(":8083", nil)
}
