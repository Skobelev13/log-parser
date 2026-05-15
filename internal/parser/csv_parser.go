package parser

import (
	"encoding/csv"
	"os"

	"log-parser/internal/model"
)

func ParseCSV(path string) ([]model.Link, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	rows, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var links []model.Link
	seen := make(map[string]bool)

	for _, row := range rows {
		if len(row) < 3 {
			continue
		}

		switchName := row[0]
		port := row[1]
		peer := row[2]

		if switchName == "NodeGUID" ||
			switchName == "NodeGuid" ||
			switchName == "NodeDesc" {
			continue
		}

		if switchName == "" || port == "" || peer == "" {
			continue
		}

		if switchName == port {
			continue
		}

		if port == "49152" || peer == "MMM-MAV" {
			continue
		}

		key := switchName + "|" + port + "|" + peer
		if seen[key] {
			continue
		}
		seen[key] = true

		link := model.Link{
			Switch: switchName,
			Port:   port,
			Peer:   peer,
		}

		links = append(links, link)
	}

	return links, nil
}
