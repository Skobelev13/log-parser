package parser

import (
	"bufio"
	"os"
	"strings"

	"log-parser/internal/model"
)

func Parse(path string) ([]model.Link, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var links []model.Link

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.Contains(line, "->") {
			continue
		}

		parts := strings.Split(line, "->")
		if len(parts) != 2 {
			continue
		}

		left := strings.TrimSpace(parts[0])
		right := strings.TrimSpace(parts[1])

		link := model.Link{
			Switch: left,
			Port:   "unknown",
			Peer:   right,
		}

		links = append(links, link)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return links, nil
}