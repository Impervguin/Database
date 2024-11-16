package task7

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
}

type Task7Storage struct {
	db *gorm.DB
}

func NewTask7Storage() (*Task7Storage, error) {
	db, err := gorm.Open(postgres.Open(GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &Task7Storage{db: db}, nil
}
