package main

import (
	"log"
	"net/http"
	"os"

	"log-parser/internal/api"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/api/v1/parse", api.ParseHandler)
	http.HandleFunc("/api/v1/csv", api.ParseCSVHandler)
	log.Printf("server started on :%s", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
