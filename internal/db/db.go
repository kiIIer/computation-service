package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type DB struct {
	*gorm.DB
}

type ComputationTask struct {
	ID        uint   `gorm:"primaryKey"`
	Status    string `gorm:"type:varchar(50)"`
	Input     string
	Result    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func NewDB() (*DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&ComputationTask{}); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
