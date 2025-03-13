package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type DB struct {
	*sql.DB
}

type ComputationTask struct {
	ID     int
	Status string
	Input  string
	Result string
}

func NewDB() (*DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Initialize schema
	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS computation_tasks (
            id SERIAL PRIMARY KEY,
            status VARCHAR(50),
            input TEXT,
            result TEXT
        )`)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
