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

	for _, row := range rows {
		if len(row) < 3 {
			continue
		}

		if row[0] == "NodeGUID" ||
			row[0] == "NodeGuid" ||
			row[0] == "NodeDesc" {
			continue
		}

		link := model.Link{
			Switch: row[0],
			Port:   row[1],
			Peer:   row[2],
		}

		links = append(links, link)
	}

	return links, nil
}
