package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var mux = http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	var port = envPortOr("8080")

	fmt.Println("starting server on port " + port[1:])
	log.Fatal(http.ListenAndServe(port, mux))
}

func envPortOr(port string) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	return ":" + port
}
