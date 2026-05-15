package parser

import (
	"bufio"
	"os"
	"strings"

	"log-parser/internal/model"
)

func ParseFile(path string) ([]model.Node, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var nodes []model.Node
	var current model.Node

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			if current.Name != "" {
				nodes = append(nodes, current)
			}

			current = model.Node{}
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		value := parts[1]

		switch key {
		case "node":
			current.Name = value

		case "ip":
			current.IP = value

		case "port":
			current.Port = value
		}
	}

	if current.Name != "" {
		nodes = append(nodes, current)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nodes, nil
}