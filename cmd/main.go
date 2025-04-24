package main

import (
	"log"
	"net/http"
	"os"

	"targetApi/internal/delivery"
	"targetApi/internal/transport"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	svc := delivery.NewService()
	handler := transport.NewHTTPHandler(svc)

	log.Printf("api listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
