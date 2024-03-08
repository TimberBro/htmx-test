package main

import (
	"log"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("This is /hello route"))
	})

	err := http.ListenAndServe(":5000", serveMux)
	if err != nil {
		log.Fatalf("Unable to start server: %+v", err)
		return
	}
}
