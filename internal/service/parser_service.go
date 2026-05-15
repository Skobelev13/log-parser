package service

import (
	"log-parser/internal/model"
	"log-parser/internal/parser"
	"log-parser/internal/storage"
)

type ParserService struct {
	linkRepo *storage.LinkRepository
}

func NewParserService(linkRepo *storage.LinkRepository) *ParserService {
	return &ParserService{linkRepo: linkRepo}
}

func (s *ParserService) Parse(path string) ([]model.Link, error) {
	return parser.Parse(path)
}

func (s *ParserService) ParseCSV(path string) ([]model.Link, error) {
	links, err := parser.ParseCSV(path)
	if err != nil {
		return nil, err
	}

	err = s.linkRepo.SaveLinks(links)
	if err != nil {
		return nil, err
	}

	return links, nil
}
