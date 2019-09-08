package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("server hello started")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	http.ListenAndServe(":8002", nil)
}
