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
	http.HandleFunc("/stage1", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "stage1")
	})

	http.ListenAndServe(":8082", nil)
}
