package api

import (
	"encoding/json"
	"net/http"

	"log-parser/internal/service"
)

type ParseRequest struct {
	Path string `json:"path"`
}

var parserService = service.NewParserService()

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	var req ParseRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	lines, err := parserService.Parse(req.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"status": "accepted",
		"path":   req.Path,
		"lines":  lines,
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(resp)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}

func ParseCSVHandler(w http.ResponseWriter, r *http.Request) {
	links, err := parserService.ParseCSV("data/ibdiagnet2.db_csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	err = encoder.Encode(links)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
