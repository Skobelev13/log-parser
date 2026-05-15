package service

import (
	"log-parser/internal/model"
	"log-parser/internal/parser"
)

type ParserService struct{}

func NewParserService() *ParserService {
	return &ParserService{}
}

func (s *ParserService) Parse(path string) ([]model.Node, error) {
	return parser.ParseFile(path)
}

func (s *ParserService) ParseCSV(path string) ([]model.Link, error) {
	return parser.ParseCSV(path)
}
