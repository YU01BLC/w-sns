package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<html><body><h1>Hello World</h1></body></html>")
	})

	port := 8080
	host := "0.0.0.0"
	addr := fmt.Sprintf("%s:%d", host, port)

	log.Printf("Server is running at http://%s\n", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
