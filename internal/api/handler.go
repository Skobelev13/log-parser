package api

import (
	"encoding/json"
	"net/http"

	"log-parser/internal/storage"
)

var linkRepo *storage.LinkRepository

func Init(repo *storage.LinkRepository) {
	linkRepo = repo
}

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status": "accepted",
		"path":   "data/test.log",
		"lines": []map[string]string{
			{
				"name": "switch-1",
				"ip":   "10.0.0.1",
				"port": "eth0",
			},
			{
				"name": "host-1",
				"ip":   "10.0.0.2",
				"port": "eth1",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ParseCSVHandler(w http.ResponseWriter, r *http.Request) {
	links, err := linkRepo.GetLinks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}

func GetLinksHandler(w http.ResponseWriter, r *http.Request) {
	links, err := linkRepo.GetLinks()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(links)
}
