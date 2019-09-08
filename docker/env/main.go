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
	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "")
	})

	http.ListenAndServe(":8081", nil)
}
