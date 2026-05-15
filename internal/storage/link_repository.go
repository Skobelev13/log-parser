package storage

import (
	"database/sql"

	"log-parser/internal/model"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) SaveLinks(links []model.Link) error {
	for _, link := range links {
		_, err := r.db.Exec(
			"INSERT INTO links (switch_name, port, peer) VALUES ($1, $2, $3)",
			link.Switch,
			link.Port,
			link.Peer,
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (r *LinkRepository) GetLinks() ([]model.Link, error) {
	rows, err := r.db.Query(
		"SELECT switch_name, port, peer FROM links",
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var links []model.Link

	for rows.Next() {
		var link model.Link

		err := rows.Scan(
			&link.Switch,
			&link.Port,
			&link.Peer,
		)
		if err != nil {
			return nil, err
		}

		links = append(links, link)
	}

	return links, nil
}
