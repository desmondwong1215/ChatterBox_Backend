package storage

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBname   string
	SSLMode  string
}

// connect to the database
func NewConnection(config *Config) (*gorm.DB, error) {

	// local database
	// dsn := fmt.Sprintf(
	// 	"host=%s port=%s password=%s user=%s dbname=%s sslmode=%s",
	// 	config.Host, config.Port, config.Password,
	// 	config.User, config.DBname, config.SSLMode,
	// )

	// deployed database
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
