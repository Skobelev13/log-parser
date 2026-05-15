package storage

import "database/sql"

func RunMigrations(db *sql.DB) error {
	query := `
CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    switch_name TEXT NOT NULL,
    port TEXT NOT NULL,
    peer TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
`
	_, err := db.Exec(query)
	return err
}
