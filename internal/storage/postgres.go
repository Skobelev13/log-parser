package storage

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://postgres:postgres@localhost:5432/logs?sslmode=disable"
	}

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	var pingErr error

	for i := 0; i < 10; i++ {
		pingErr = db.Ping()
		if pingErr == nil {
			return db, nil
		}

		time.Sleep(2 * time.Second)
	}

	return nil, pingErr
}
