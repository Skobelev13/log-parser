package main

import (
	"log"
	"net/http"
	"os"

	"log-parser/internal/api"

	"log-parser/internal/storage"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db, err := storage.NewPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = storage.RunMigrations(db)
	if err != nil {
		log.Fatal(err)
	}

	linkRepo := storage.NewLinkRepository(db)
	api.Init(linkRepo)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/api/v1/parse", api.ParseHandler)
	http.HandleFunc("/api/v1/csv", api.ParseCSVHandler)
	http.HandleFunc("/api/v1/links", api.GetLinksHandler)

	log.Printf("server started on :%s", port)

	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}

}
