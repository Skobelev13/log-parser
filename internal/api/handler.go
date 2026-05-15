package api

import (
	"encoding/json"
	"net/http"

	"log-parser/internal/parser"
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
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func ParseCSVHandler(w http.ResponseWriter, r *http.Request) {
	links, err := parser.ParseCSV("data/ibdiagnet2.db_csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = linkRepo.SaveLinks(links)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": "saved",
		"count":  len(links),
		"links":  links,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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
